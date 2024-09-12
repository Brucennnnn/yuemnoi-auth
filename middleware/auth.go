package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		// If not, redirect to the login page
		log.Println("Auth middleware")
		next.ServeHTTP(w, r)
	})
}
