package token

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

func SetTokenToCookie(ctx context.Context, req GenerateTokenRequest) (string, error) {
	tokens, err := generateTokens(req)
	if err != nil {
		return "", err
	}

	err = setCookieInResponse(ctx, tokens.Refresh())
	if err != nil {
		return "", err
	}

	return tokens.Refresh(), nil
}

// setCookieInResponse sets the authentication cookie
func setCookieInResponse(ctx context.Context, token string) error {
	md := metadata.New(
		map[string]string{
			"set-cookie": fmt.Sprintf(
				"%s=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=Lax",
				RefreshCookie, token, 60*60*24*60,
			),
		},
	)

	return grpc.SetHeader(ctx, md)
}

func generateTokens(req GenerateTokenRequest) (TokenPair, error) {
	accessExp := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(req.JwtKey))
	if err != nil {
		return TokenPair{}, err
	}

	refreshExp := time.Now().Add(14 * 24 * time.Hour)
	refreshClaims := &Claims{
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(req.RefreshKey))
	if err != nil {
		return TokenPair{}, err
	}

	return NewTokenPair(accessString, refreshString), nil
}

// RefreshClaimsFromContext извлекает и валидирует refresh token из gRPC контекста
func RefreshClaimsFromContext(ctx context.Context, refreshSecret string) (*Claims, error) {
	// Получаем метаданные из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata not found")
	}

	// Ищем cookie в метаданных
	cookieHeader := md.Get(RefreshCookie)
	if len(cookieHeader) == 0 {
		return nil, fmt.Errorf("cookie header not found")
	}

	// Валидируем JWT токен
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookieHeader[0], claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(refreshSecret), nil
	})

	if err != nil || !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func AccessClaimsFromRequest(authHeader, jwtAccessSecret string) (*Claims, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("invalid token")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token")
	}
	tokenStr := parts[1]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtAccessSecret), nil
	})
	if err != nil || !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
