package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type CorsConfig interface {
	GetAllowedHosts() []string
}

func CORSMiddleware(corsConfig CorsConfig) func(next http.Handler) http.Handler {
	allowedHosts := make(map[string]struct{})
	for _, host := range corsConfig.GetAllowedHosts() {
		allowedHosts[host] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowedOrigin := ""
			hasValidAPIKey := requestHasValidAPIKey(r)
			decision := "pending"

			if len(origin) != 0 {
				if os.Getenv("DEBUG") == "true" {
					allowedOrigin = origin
					decision = "origin allowed by DEBUG"
				} else if hasValidAPIKey {
					allowedOrigin = origin
					decision = "origin allowed by api_key"
				} else if u, err := url.Parse(origin); err == nil {
					if _, exists := allowedHosts[u.Host]; exists {
						allowedOrigin = origin
						decision = "origin allowed by allowed_hosts"
					} else {
						decision = "origin rejected: host not in allowed_hosts"
					}
				} else {
					decision = "origin rejected: parse error"
				}
			}

			if len(origin) == 0 && !hasValidAPIKey {
				fmt.Printf(
					"[cors] method=%s path=%s host=%s origin=%q content_type=%q api_key=%t allowed_origin=%q decision=%s status=%d\n",
					r.Method,
					r.URL.Path,
					r.Host,
					origin,
					r.Header.Get("Content-Type"),
					hasValidAPIKey,
					allowedOrigin,
					"rejected: empty origin and invalid api_key",
					http.StatusForbidden,
				)
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if hasValidAPIKey && len(origin) == 0 {
				allowedOrigin = "getcourse.ru"
				decision = "empty origin allowed by api_key"
			}

			if len(allowedOrigin) != 0 {
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == "OPTIONS" {
				if len(origin) != 0 && len(allowedOrigin) == 0 {
					fmt.Printf(
						"[cors] method=%s path=%s host=%s origin=%q content_type=%q api_key=%t allowed_origin=%q decision=%s status=%d\n",
						r.Method,
						r.URL.Path,
						r.Host,
						origin,
						r.Header.Get("Content-Type"),
						hasValidAPIKey,
						allowedOrigin,
						decision,
						http.StatusForbidden,
					)
					w.WriteHeader(http.StatusForbidden)
					return
				}
				fmt.Printf(
					"[cors] method=%s path=%s host=%s origin=%q content_type=%q api_key=%t allowed_origin=%q decision=%s status=%d\n",
					r.Method,
					r.URL.Path,
					r.Host,
					origin,
					r.Header.Get("Content-Type"),
					hasValidAPIKey,
					allowedOrigin,
					decision,
					http.StatusNoContent,
				)
				w.WriteHeader(http.StatusNoContent)
				return
			}

			fmt.Printf(
				"[cors] method=%s path=%s host=%s origin=%q content_type=%q api_key=%t allowed_origin=%q decision=%s status=pass\n",
				r.Method,
				r.URL.Path,
				r.Host,
				origin,
				r.Header.Get("Content-Type"),
				hasValidAPIKey,
				allowedOrigin,
				decision,
			)

			next.ServeHTTP(w, r)
		})
	}
}

func requestHasValidAPIKey(r *http.Request) bool {
	apiKey := os.Getenv("GK_SUBS_API_KEY")
	if apiKey == "" || r.Body == nil || r.Method == http.MethodOptions {
		return false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		values, err := url.ParseQuery(string(body))
		if err != nil {
			return false
		}
		return values.Get("api_key") == apiKey
	}

	if strings.Contains(contentType, "application/json") {
		payload := map[string]interface{}{}
		if err := json.Unmarshal(body, &payload); err != nil {
			return false
		}

		value, ok := payload["api_key"].(string)
		return ok && value == apiKey
	}

	return false
}

//import (
//"github.com/rs/cors"
//)
//
//// Настраиваем CORS
//c := cors.New(cors.Options{
//AllowedOrigins:   []string{"https://example.com", "https://api.example.com", "http://localhost:3000"},
//AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//AllowedHeaders:   []string{"Content-Type", "Authorization"},
//AllowCredentials: true,
//Debug:           true, // Только для разработки
//})
