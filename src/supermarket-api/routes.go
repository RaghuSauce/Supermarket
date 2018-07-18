package supermarket_api

import (
	"net/http"
	"SupermarketAPI/src/supermarket-service"
)

// struct of required fields for mux router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Array of all routes in application
type Routes []Route


var routes = Routes{
	//Index Greeting page
	Route{
		"Index",
		"GET",
		"/",
		supermarket_service.Index,
	},

	//Get the list of all produce currently available
	Route{
		"Fetch",
		"GET",
		"/fetch",
		supermarket_service.FetchProduceList,
	},

	//Add a produce Item
	Route{
		"Add",
		"POST",
		"/add",
		supermarket_service.AddProduceItem,
	},

	//Remove A produce Item
	Route{
		"Remove",
		"Delete",
		"/remove",
		supermarket_service.RemoveProduceItem,
	},
}
