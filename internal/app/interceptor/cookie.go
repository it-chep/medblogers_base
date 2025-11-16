package interceptor

import (
	"context"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func CookieInterceptor(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	//if loginResp, ok := resp.(*desc.LoginResponse); ok {
	//	// Устанавливаем refresh token cookie
	//	http.SetCookie(w, &http.Cookie{
	//		Name:     "refresh_token",
	//		Value:    loginResp.GetRefreshToken(),
	//		Expires:  time.Now().Add(60 * 24 * time.Hour),
	//		HttpOnly: true,
	//		Path:     "/",
	//	})
	//
	//	// Очищаем refresh token из тела ответа для безопасности
	//	loginResp.RefreshToken = ""
	//}
	return nil
}
