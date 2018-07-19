package main

import (
	"SupermarketAPI/supermarket_api"
	"log"
	"net/http"
	"fmt"
)

func main() {
	var ver float32 = 0.00
	fmt.Println("ver",ver)
	router := supermarket_api.Supermarket_router()
	log.Fatal(http.ListenAndServe(":8080", router))

	//item := supermarket_database.ProduceItem{
	//	ProduceCode: "A12T-4GH7-QPL9-3N4M",
	//	Name:        "Lettuce",
	//	UnitPrice:   "300",
	//}
	//fmt.Println(supermarket_database.ValidateProduceItem(item))
}
