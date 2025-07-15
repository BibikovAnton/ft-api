package middleware

import (
	"fmt"
	"net/http"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		next.ServeHTTP(w, r)
	})
}
