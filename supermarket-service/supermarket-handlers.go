package supermarket_service

import (
	"net/http"
	"fmt"
	"html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "supermarketApi, %q", html.EscapeString(r.URL.Path))
}