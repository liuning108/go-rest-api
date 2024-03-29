package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		//Bearer: token-string
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "bad auth header", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accessToken string) bool {

	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not handle this token")
		}

		return mySigningKey, nil

	})
	if err != nil {
		return false
	}

	return token.Valid
}
