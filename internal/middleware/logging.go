package middleware

import (
	"html"
	"log"
	"net/http"
)

func LogRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, html.EscapeString(r.URL.Path))
}
