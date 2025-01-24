package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// AuthMiddleware checks for valid JWT token in Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// Validate token (for simplicity, just check if it starts with "Bearer")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Assuming we have a function `validateJWT` to check the token validity
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !validateJWT(token) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Continue if valid
		next.ServeHTTP(w, r)
	})
}

// validateJWT is a mock function, replace it with actual JWT verification logic.
func validateJWT(token string) bool {
	// Implement your JWT validation logic here
	// For example: check signature, expiration, etc.
	fmt.Println("Validating token:", token)
	return true // Just a placeholder
}
