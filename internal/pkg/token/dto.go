package token

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshCookie = "medblogers_refresh"
)

type GenerateTokenRequest struct {
	Email      string
	JwtKey     string
	RefreshKey string
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string
	refreshToken string
}

func NewTokenPair(accessToken, refreshToken string) TokenPair {
	return TokenPair{
		AccessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (p TokenPair) Refresh() string {
	return p.refreshToken
}
