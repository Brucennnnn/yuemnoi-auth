package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
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

	return c.JSON(userInfo)
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
