package supermarket_service

import (
	"SupermarketAPI/supermarket_database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
		"net/http"
	"strings"
)

//Get Mapping	"/"
func Index(w http.ResponseWriter, r *http.Request) {
	//file, err := ioutil.ReadFile("VERSION")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Print()
	//string(file)
	fmt.Fprintf(w, "%s%s", "Supermarket-API:", "0.0.1")
}

//Get Mapping  "/fetch "
func FetchProduceList(w http.ResponseWriter, r *http.Request) {
	c := make(chan []supermarket_database.ProduceItem)
	go supermarket_database.ListProduceItems(c)
	db := <-c
	json.NewEncoder(w).Encode(db)
}

//Post Mapping	"/add"
//TODO Handle Invalid JSON Requests
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

	if isValid, errs := supermarket_database.ValidateProduceItem(produce); err == nil && isValid {
		if e := supermarket_database.AddProduceItemToDatabase(produce); e == nil {
			fmt.Fprint(w, "Success")
		} else {
			fmt.Fprint(w, e)
		}
	} else {
		var errorString string
		for _, err := range errs {
			errorString += err.Error() + "\n"
		}
		fmt.Fprint(w, "Produce Item is invalid for the following reasons \n\n", errorString)
		//fmt.Fprint(w, err)
	}

}

//Delete Mapping	"/remove"
func RemoveProduceItem(w http.ResponseWriter, r *http.Request) {
	produceCode := getProduceCodeUrlParamter(r)
	if err := supermarket_database.RemoveProduceItemFromDatabase(produceCode); err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, "Success")
	}

}

func getProduceCodeUrlParamter(r *http.Request) string {
	vars := mux.Vars(r) //Get url variables
	code := strings.Split(vars["code"], "=")
	return code[1]
}
