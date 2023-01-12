package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(
	original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "no authorizado", http.StatusUnauthorized)
			return
		}

		// Bearer token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "no autorizado", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "no authorizado", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("misionimposible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("no pudo validarse el token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
