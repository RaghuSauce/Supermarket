package main

import (
	"SupermarketAPI/supermarket_api"
	"log"
	"net/http"
)

func main() {
	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8081", router))

}
