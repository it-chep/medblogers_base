package middleware

import (
	"net/http"
)

// todo сделать нормальный cors
var allowedOrigins = map[string]bool{
	"https://example.com":     true,
	"https://api.example.com": true,
	"http://localhost:3000":   true, // для разработки
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Проверяем, есть ли домен в списке разрешенных
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Для предварительных OPTIONS запросов
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
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
