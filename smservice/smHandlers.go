package smservice

import (
	"SupermarketChallenge/smdb"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"os"
)

const(
	BODYSIZELIMT = 1048576		// max payload size for posts
)

//Get Mapping	"/"
func Index(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("VERSION") // read the version file
	if err != nil {                         //if an error occurs tell the user
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	} else { // else return the version of the api in use
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "%s%s", "Supermarket-API:", string(file))
	}
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	produceCode := getProduceCodeUrlParameter(r)        //get the code from the request
	fileExits, item := smdb.GetProduceItem(produceCode) //see if it exits

	if fileExits { //if it does return it
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
		json.NewEncoder(w).Encode(item)
	} else { // else let the user know it wasn't found
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "produce item with code %s not found in the database", produceCode)
	}
}

//Get Mapping  "/fetch "
func FetchProduceList(w http.ResponseWriter, r *http.Request) {
	c := make(chan []smdb.ProduceItem)                                 //make for the list of produce items
	defer close(c)
	go smdb.ListProduceItems(c)                                        //populate the channel of the items
	db := <-c                                                          //get the items in the channel
	w.WriteHeader(http.StatusAccepted)                                                 //set the response code
	w.Header().Set("Content-Type", "application/json ; charset=UTF-8") //Set the response type
	json.NewEncoder(w).Encode(db)                                      //return the list
}

//Get Mapping "/Logs"
func GetLogs(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("rest.log") // just pass the file name
	if err != nil {                       //if there is an error return to the user and set header to 500
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	str := string(b)     // convert content to a 'string'
	w.Header().Set("Content-Type", "application/json ; charset=UTF-8") //Set the response type
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, str) // print the content as a 'string'
}

//Post Mapping	"/add"
func AddProduceItem(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Add, %q", html.EscapeString(r.URL.Path))
	var produce smdb.ProduceItem // Declare a produce Item to to unmarshal into

	//w.Header().Set("Content-Type", "application/json ; charset=UTF-8") //Set the response type
	w.WriteHeader(http.StatusInternalServerError)                          //Set the response Code

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, BODYSIZELIMT)) // Read the body of the request and limit the body size to 1MB
	if err != nil {
		fmt.Fprintln(w,err)
	}
	if json.Unmarshal(body, &produce); err != nil { //Unmarshal the request into the struct, panic if an error occurs

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Fprintln(w,err)
		}
	}

	if isValid, errs := smdb.ValidateProduceItem(produce); err == nil && isValid {
		if e := smdb.AddProduceItemToDatabase(produce); e == nil {
			w.WriteHeader(http.StatusAccepted)		//Set the response Code
			w.Header().Set("Content-Type", "application/json ; charset=UTF-8") //Set the response type
			fmt.Fprint(w, "Success")
		} else {
			fmt.Fprint(w, e)
		}
	} else {
		var errorString string
		for _, err := range errs {
			errorString += err.Error() + "\n"
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Produce Item is invalid for the following reasons \n\n", errorString)
	}
}

//Delete Mapping	"/remove"
//Attempt to remove an item from the database
func RemoveProduceItem(w http.ResponseWriter, r *http.Request) {
	produceCode := getProduceCodeUrlParameter(r)
	if err := smdb.RemoveProduceItemFromDatabase(produceCode); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "Success")
	}
}

//Delete Mapping 		"/purgeLogs"
//Wipe the logs out
func CleanLogs(w http.ResponseWriter, r *http.Request) {
	if err := os.Remove("rest.log"); err != nil { //remove the file containing the logs
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "Logs Purged")
	}
}

//func to get package name from url
func getProduceCodeUrlParameter(r *http.Request) string {
	vars := mux.Vars(r) //Get url variables
	code := strings.Split(vars["code"], "=")
	return code[1]
}
