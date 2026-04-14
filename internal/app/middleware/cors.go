package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
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

			fmt.Printf("ORIGIN: %s\n", origin)

			if origin != "" {
				if os.Getenv("DEBUG") == "true" {
					allowedOrigin = origin
				} else if u, err := url.Parse(origin); err == nil {
					if _, exists := allowedHosts[u.Host]; exists {
						allowedOrigin = origin
					}
				}
			}

			if allowedOrigin != "" {
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == "OPTIONS" {
				if origin != "" && allowedOrigin == "" {
					w.WriteHeader(http.StatusForbidden)
					return
				}
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
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
