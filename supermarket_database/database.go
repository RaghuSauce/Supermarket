package supermarket_database

import (
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var database = []ProduceItem{}

//var database = map[string]ProduceItem{}

func init() {
	//Array representing the initial state of the database
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

//Returns the all of the ProduceItems values from the database as a slice
func ListProduceItems() []ProduceItem {
	var db []ProduceItem
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		db = database
	}()
	wg.Wait()
	return db
}

//All produce items are assumed that they will enter via the api, thus validation will occur at the api layer
func AddProduceItemToDatabase(item ProduceItem) error {
	var err error // init error for goroutine to set
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := validateUUID(item.ProduceCode); e == nil {
			database = append(database, item)
			err = nil
		} else {
			err = errors.New("Error Adding Produce Item to the Database")
		}
	}()
	wg.Wait()
	return err
}

func RemoveProduceItemFromDatabase(produceCode string) error {
	var err error //initial error for the go routine to modify
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		if e := validateUUID(produceCode); e != nil {
			newDatabase := []ProduceItem{}
			for _, element := range database {
				if element.ProduceCode != produceCode { //if the produce codes are equal
					// a = append(a[:i], a[i+1:]...)	some research unveiled this, might be more performant but I don't fully understand it yet
					newDatabase = append(newDatabase, element)
				}
			}
			database = newDatabase
			err = nil
		} else {
			err = errors.New("Error Removing Produce Item from the Database")
		}
	}()

	wg.Wait()
	return err
}

/*Checks to see if the Produce code already exists,
if yes, then returns a message
else returns a nil error

does not need to be synced bc it it is always called from a synced routine
*/
func validateUUID(produceCode string) error {
	for _, element := range database {
		if strings.ToUpper(element.ProduceCode) == strings.ToUpper(produceCode) {
			return errors.New("Produce with this Code already exists")
		}
	}
	return nil
}
