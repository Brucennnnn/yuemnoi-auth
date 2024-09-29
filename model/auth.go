package model

import "github.com/golang-jwt/jwt/v5"

type AuthTokenClaim struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
