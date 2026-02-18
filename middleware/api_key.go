package middleware

import "net/http"

// func (api key) -> func handler http.handler
// APIKey middleware untuk validate API key
func APIKey(validAPIKey string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			if apiKey == "" {
				http.Error(w, "API Key required", http.StatusUnauthorized)
				return
			}

			if apiKey != validAPIKey {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			// API key valid, continue
			next(w, r)
		}
	}
}
