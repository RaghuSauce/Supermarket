package supermarket_api

import "net/http"

// struct of required fields for mux router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Array of all routes in application
type Routes []Route

//TODO add handler functions
var routes = Routes{
	//Index Greeting page
	Route{
		"Index",
		"GET",
		"/",
		nil,
	},

	//Get the list of all produce currently available
	Route{
		"Fetch",
		"GET",
		"/fetch",
		nil,
	},

	//Add a produce Item
	Route{
		"Add",
		"POST",
		"/add",
		nil,
	},

	//Remove A produce Item
	Route{
		"Remove",
		"Delete",
		"/remove",
		nil,
	},
}
