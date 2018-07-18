package supermarket_service

import (
	"net/http"
	"fmt"
	"html"
)
//TODO Implement Handler Stubs

//Get Mapping
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index, %q", html.EscapeString(r.URL.Path))
}

//Get Mapping
func FetchProduceList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fetch, %q", html.EscapeString(r.URL.Path))
}

//Post Mapping
func AddProduceItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add, %q", html.EscapeString(r.URL.Path))
}

//Delete Mapping
func RemoveProduceItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove, %q", html.EscapeString(r.URL.Path))
}
