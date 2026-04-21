package middleware

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		godotenv.Load()
		apiKey := os.Getenv("API_KEY")

		key := r.Header.Get("X-API-Key")
		if key != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
