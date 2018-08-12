package smapi

import (
	"SupermarketChallenge/smservice"
	"net/http"
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
		smservice.Index,
	},

	//Get the list of all produce currently available
	Route{
		"Fetch",
		"GET",
		"/fetch",
		smservice.FetchProduceList,
	},

	//Get the list of all produce currently available
	Route{
		"GetOne",
		"GET",
		"/get/{code}",
		smservice.GetOne,
	},

	//Get the stored logs
	Route{
		"GetLogs",
		"GET",
		"/logs",
		smservice.GetLogs,
	},

	//Add a produce Item
	Route{
		"Add",
		"POST",
		"/add",
		smservice.AddProduceItem,
	},

	//Remove A produce Item
	Route{
		"Remove",
		"DELETE",
		"/remove/{code}",
		smservice.RemoveProduceItem,
	},
	//Purge the log
	Route{
		"CleanLogs",
		"DELETE",
		"/purgeLogs",
		smservice.CleanLogs,
	},
}
