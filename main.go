package main

import (
	"SupermarketChallenge/smapi"
	"log"
	"net/http"
)

func main() {
	router := smapi.SupermarketRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
