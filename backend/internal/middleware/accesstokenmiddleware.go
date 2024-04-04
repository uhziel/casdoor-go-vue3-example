package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	errInvalidToken = fmt.Errorf("invalid token")
)

type AccessTokenMiddleware struct {
	publicKey *rsa.PublicKey
}

func NewAccessTokenMiddleware(publicKey *rsa.PublicKey) *AccessTokenMiddleware {
	return &AccessTokenMiddleware{
		publicKey: publicKey,
	}
}

func (m *AccessTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return m.publicKey, nil
		})

		if err != nil {
			logx.Errorf("AccessTokenMiddleware authorize failed: %w", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logx.Errorf("AccessTokenMiddleware authorize failed: %w", errInvalidToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logx.Errorf("AccessTokenMiddleware claims invalid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			if k == "name" {
				ctx = context.WithValue(ctx, k, v)
			}
		}

		next(w, r.WithContext(ctx))
	}
}
