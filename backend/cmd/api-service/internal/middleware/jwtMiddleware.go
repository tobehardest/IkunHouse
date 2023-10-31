package middleware

import (
	"net/http"
	"strings"
	"video_clip/cmd/api-service/internal/config"
	"video_clip/pkg/jwtx"
)

type JWTMiddleware struct {
	Config config.Config
}

func NewJWTMiddleware(c config.Config) *JWTMiddleware {
	return &JWTMiddleware{
		Config: c,
	}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var token string

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			//httpx.OkJson(w,)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			//httpx.OkJson()
			return
		}

		token = parts[1]

		isTimeout, err := jwtx.ValidToken(token, m.Config.Auth.AccessSecret)
		if err != nil || isTimeout {
			//httpx.OkJson()
			return
		}
		next(w, r)
	}
}
