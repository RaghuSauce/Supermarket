package supermarket_service

import (
	"net/http"
	"fmt"
	"html"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index, %q", html.EscapeString(r.URL.Path))
}

//TODO Implement Handler Stubs
func FetchProduceList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fetch, %q", html.EscapeString(r.URL.Path))
}

func AddProduceItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add, %q", html.EscapeString(r.URL.Path))
}

func RemoveProduceItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove, %q", html.EscapeString(r.URL.Path))
}