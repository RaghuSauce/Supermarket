package main

import (
	"SupermarketAPI/supermarket_api"
	"log"
	"net/http"
)

func main() {
	router := supermarket_api.SupermarketRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
