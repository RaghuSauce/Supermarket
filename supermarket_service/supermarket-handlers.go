package supermarket_service

import (
	"SupermarketAPI/supermarket_database"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
)

//TODO Implement Handler Stubs

//Get Mapping	"/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index, %q", html.EscapeString(r.URL.Path))
}

//Get Mapping  "/fetch "
func FetchProduceList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(supermarket_database.ListProduceItems())
}

//Post Mapping	"/add"
func AddProduceItem(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Add, %q", html.EscapeString(r.URL.Path))
	var produce supermarket_database.ProduceItem // Declare a produce Item to to unmarshal into

	body, err := ioutil.ReadAll(
		io.LimitReader(r.Body, 1048576)) // Read the body of the request and limit the body size to 1MB
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(body, &produce); err != nil { //Unmarshal the request into the struct, panic if an error occurs
		w.Header().Set("Content-Type", "application/json ; charset=UTF-8") //Set the response type
		w.WriteHeader(422)                                                 //Set the response Code

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//Validate the struct, and Change the case of the case insensitive fields
	if err := supermarket_database.ValidateProduceItem(&produce); err != nil {
		//TODO Actually add to the database
		//AddProduceItem(produce)
	}
	fmt.Fprint(w, produce)
}

//Delete Mapping	"/remove"
func RemoveProduceItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove, %q", html.EscapeString(r.URL.Path))
}
