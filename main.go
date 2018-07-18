package main

import (
	"SupermarketAPI/supermarket-api"
	"log"
	"net/http"
)

func main() {
	//
	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8080", router))

}

