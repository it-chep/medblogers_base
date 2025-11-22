package token

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
)

func SetTokenToCookie(ctx context.Context, req GenerateTokenRequest) error {
	tokens, err := generateTokens(req)
	if err != nil {
		return err
	}

	refreshCookie := &http.Cookie{
		Name:     RefreshCookie,
		Value:    tokens.Refresh(),
		Expires:  time.Now().UTC().Add(60 * 24 * time.Hour), // 60 дней
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // раскомментируйте для HTTPS
	}

	// Устанавливаем заголовки через gRPC метаданные
	if err = grpc.SetHeader(ctx, metadata.Pairs(
		"Set-Cookie", refreshCookie.String(),
	)); err != nil {
		return status.Error(codes.Internal, "Failed to set cookie")
	}
	return nil
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
	cookieHeader := md.Get("cookie")
	if len(cookieHeader) == 0 {
		return nil, fmt.Errorf("cookie header not found")
	}

	// Парсим cookie
	refreshToken, err := extractRefreshTokenFromCookie(cookieHeader[0])
	if err != nil {
		return nil, err
	}

	// Валидируем JWT токен
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
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

// extractRefreshTokenFromCookie извлекает refresh token из cookie строки
func extractRefreshTokenFromCookie(cookieHeader string) (string, error) {
	cookies := strings.Split(cookieHeader, ";")
	for _, cookie := range cookies {
		cookie = strings.TrimSpace(cookie)
		if strings.HasPrefix(cookie, RefreshCookie+"=") {
			token := strings.TrimPrefix(cookie, RefreshCookie+"=")
			if token == "" {
				return "", fmt.Errorf("refresh token is empty")
			}
			return token, nil
		}
	}
	return "", fmt.Errorf("refresh token cookie not found")
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
