package supermarket_database

import (
	"github.com/pkg/errors"
	"strings"
	"sync"
)

// array of produce item that is acting as the database
var database = []ProduceItem{}

//init function to set the starting state of the database
func init() {
	database = []ProduceItem{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		ProduceItem{
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			Name:        "Peach",
			UnitPrice:   "2.99",
		},
		ProduceItem{
			ProduceCode: "YRT6-72AS-K736-L4AR",
			Name:        "Green Pepper",
			UnitPrice:   "0.79",
		},
		ProduceItem{
			ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
			Name:        "Gala Apple",
			UnitPrice:   "3.59",
		},
	}
}

//Sends the current database slice via channel to caller
func ListProduceItems(c chan []ProduceItem){
	c <- database
	close(c)
}

//All produce items are assumed that they will enter via the api, thus validation will occur at the api layer
//Adds a produce item to the database if its Produce Code is valid
func AddProduceItemToDatabase(item ProduceItem) error {
	var err error         // init error for goroutine to set
	var wg sync.WaitGroup //wait group for anon async func

	wg.Add(1) //wait for func to end
	go func() {
		defer wg.Done() //decrement wait counter at end of goroutine

		c := make(chan error)                //receive error from async call from uuid check
		go validateUUID(item.ProduceCode, c) //check if produce code is valid

		if e := <-c; e == nil { // if the error from validateUUID is nil append item to the database
			database = append(database, item)
			err = nil //bc we were able to append to the database set the set error to null
		} else { //if the ProduceCode from item matched an existing item then set a non nil error
			err = errors.New("Error Adding Produce Item to the Database")
		}
	}()
	wg.Wait()  // wait for async add to finish
	return err //return error set by go routine
}

func RemoveProduceItemFromDatabase(produceCode string) error {
	var err error         //initial error for the go routine to modify
	var wg sync.WaitGroup //wait group for anon async func

	wg.Add(1) //add wait to counter for go routine

	go func() {
		defer wg.Done()                 //decrement from wait counter when routine is finished
		c := make(chan error)           //make a channel to receive an error from uuid validation
		go validateUUID(produceCode, c) //validate produceCode
		if e := <-c; e != nil {         //if the error is NOT nil i.e. the item exists within the database then remove it
			newDatabase := []ProduceItem{}     //new slice
			for _, element := range database { //iterate over existing db and add all but selected code
				if element.ProduceCode != produceCode { //if the produce codes are equal
					// a = append(a[:i], a[i+1:]...)	some research unveiled this, might be more performant but I don't fully understand it yet
					newDatabase = append(newDatabase, element) //append element to new db
				}
			}
			database = newDatabase //set db equal to new db
			err = nil              // bc remove operation was a success set the error to nil
		} else { //else produce code does not exist within the db
			err = errors.New("Error Removing Produce Item from the Database") //set error message for return
		}
	}()

	wg.Wait()  //wait for routine to finish
	return err //return what the routine set the  error as
}

/*Checks to see if the Produce code already exists,
if yes, then returns a message
else returns a nil error
*/
func validateUUID(produceCode string, err chan error) {
	var e error
	for _, element := range database {
		if strings.ToUpper(element.ProduceCode) == strings.ToUpper(produceCode) {
			e = errors.New("Produce with this Code already exists")
		}
	}
	err <- e
	close(err)
}

//Method to reset the database to initial state, used for testing
func resetDB(){
	database = []ProduceItem{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		ProduceItem{
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			Name:        "Peach",
			UnitPrice:   "2.99",
		},
		ProduceItem{
			ProduceCode: "YRT6-72AS-K736-L4AR",
			Name:        "Green Pepper",
			UnitPrice:   "0.79",
		},
		ProduceItem{
			ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
			Name:        "Gala Apple",
			UnitPrice:   "3.59",
		},
	}
}