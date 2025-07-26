package middleware

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

func MetadataMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := metadata.New(map[string]string{
			"x-forwarded-method":  r.Method,
			"x-forwarded-path":    r.URL.Path,
			"x-forwarded-address": r.RemoteAddr,
		})

		for k, v := range r.Header {
			md.Set(strings.ToLower(k), v...)
		}

		ctx := metadata.NewIncomingContext(r.Context(), md)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MetadataAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)

	for k, v := range r.Header {
		md.Set(strings.ToLower(k), v...)
	}

	md.Set("x-http-method", r.Method)
	md.Set("x-http-path", r.URL.Path)

	return md
}
