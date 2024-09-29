package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sds-2/config"
	"github.com/sds-2/feature/user"
	"github.com/sds-2/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewAuthHandler(cfg *config.Config, userRepo user.UserRepository) *AuthHandler {
	return &AuthHandler{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (h *AuthHandler) OAuthLogin(c *fiber.Ctx) error {
	oauth2Config := NewOAuthConfig(h.cfg)
	url := oauth2Config.AuthCodeURL(h.cfg.GoogleOAuthConfig.OauthStateString)
	return c.Redirect(url, 302)
}

func (h *AuthHandler) OAuthCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != h.cfg.GoogleOAuthConfig.OauthStateString {
		return c.Redirect("/", 302)
	}

	code := c.Query("code")
	oauth2Config := NewOAuthConfig(h.cfg)
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error while exchanging token:", err)
		return c.Redirect("/", 302)
	}

	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Println("Error while getting user info:", err)
		return c.Redirect("/", 302)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)

	email := userInfo["email"].(string)
	if !strings.HasSuffix(email, "@student.chula.ac.th") {
		return c.SendString("Access denied: invalid email domain")
	}

	user, err := h.userRepo.GetUserByEmail(email)
	if err != nil {
		return c.SendString("Internal server error: can not get user by email on auth handler")
	}
	if user == nil {
		user = &model.User{
			Email:   email,
			Name:    userInfo["given_name"].(string),
			SurName: userInfo["family_name"].(string),
		}
		_, err := h.userRepo.CreateUser(*user)
		if err != nil {
			return c.SendString("Internal server error: can not create user on auth handler")
		}
	}

	cookieToken, err := createToken(user.ID, h.cfg)
	if err != nil {
		return c.SendString("Internal server error: failed to create token")

	}
	cookie := new(fiber.Cookie)
	cookie.Name = h.cfg.Cookie.CookieNameAuth
	cookie.Value = cookieToken
	cookie.Expires = time.Now().Add(h.cfg.Cookie.Expires) // Set cookie expiry
	cookie.HTTPOnly = true                                // Make the cookie inaccessible to JavaScript
	cookie.Secure = true
	if h.cfg.Environment == "dev" {
		cookie.Secure = false
	}
	cookie.SameSite = "Strict" // Prevent CSRF

	c.Cookie(cookie)
	return c.Redirect("/", 302)
}

func createToken(userID int, config *config.Config) (string, error) {
	// Create custom claims with user data and standard claims
	claims := model.AuthTokenClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token expiration
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // Issued at
			Issuer:    config.AppName,                                    // Issuer
		},
	}

	// Create token using signing method HS256 and the custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(config.Cookie.Secret)

	// Sign and return the token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GoogleOAuthConfig.ClientID,
		ClientSecret: cfg.GoogleOAuthConfig.ClientSecret,
		RedirectURL:  cfg.GoogleOAuthConfig.RedirectURL,
		Scopes:       cfg.GoogleOAuthConfig.Scopes,
		Endpoint:     google.Endpoint,
	}
}
