package main

import (
	"SupermarketAPI/supermarket_api"
	"log"
	"net/http"
)

func main() {

	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8080", router))

	//fmt.Println(supermarket_database.ListProduceItems())
	//fmt.Println(supermarket_database.ValidateUUID("A12T-4GH7-QPL9-3N4M"))
	//fmt.Println(supermarket_database.ValidateUUID("12"))

}
