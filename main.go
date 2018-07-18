package main

import (
	"log"
	"net/http"

	"SupermarketAPI/supermarket-api"
)

func main() {

	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8080", router))
}

