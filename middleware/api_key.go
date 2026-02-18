package middleware

import (
	"net/http"
)

// func (api key) -> func handler http.handler
// APIKey middleware untuk validate API key
func APIKey(validApiKey string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			// Debug logging
			// fmt.Printf("=== API Key Debug ===\n")
			// fmt.Printf("Expected API Key: '%s' (length: %d)\n", validApiKey, len(validApiKey))
			// fmt.Printf("Received API Key: '%s' (length: %d)\n", apiKey, len(apiKey))
			// fmt.Printf("Match: %v\n", apiKey == validApiKey)
			// fmt.Printf("====================\n")

			if apiKey == "" {
				http.Error(w, "API Key required", http.StatusUnauthorized)
				return
			}

			if apiKey != validApiKey {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			// API key valid, continue
			next(w, r)
		}
	}
}
