package middleware

import (
	"net/http"
	"log"
)

// LogRequest logs incoming HTTP requests
func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
