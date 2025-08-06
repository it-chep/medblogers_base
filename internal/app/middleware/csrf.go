package middleware

// todo сделать нормальный CSRF
//func CSRFMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		csrfMiddleware := csrf.Protect(
//			[]byte("32-байтный-секретный-ключ-здесь"),
//			csrf.Secure(true),
//			csrf.Path("/"),
//			csrf.HttpOnly(true),
//			csrf.SameSite(csrf.SameSiteStrictMode),
//		)
//
//		return csrfMiddleware(next)
//	})
//}
