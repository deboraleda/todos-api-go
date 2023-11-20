package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			http.Error(w, `{"message":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		reqToken := strings.Split(bearerToken, " ")[1]
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		log.Printf("%v - %v ", tkn, claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, `{"message":"unauthorized"}`, http.StatusUnauthorized)
				return
			}
			http.Error(w, `{"message":"bad request"}`, http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, `{"message":"Invalid token"}`, http.StatusForbidden)
			return
		}
		log.Printf("decoded completed!")

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "userId", claims.Id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
