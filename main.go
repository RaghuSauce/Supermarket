package main

import (
	"fmt"
	"SupermarketAPI/supermarket_api"
	"log"
	"net/http"
)

func main() {
	var ver float32 = 0.00
	fmt.Println("ver",ver)
	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
