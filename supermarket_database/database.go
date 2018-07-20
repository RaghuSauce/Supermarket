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
func ListProduceItems(c chan  []ProduceItem) {
	c <- database
	close(c)
}

//All produce items are assumed that they will enter via the api, thus validation will occur at the api layer
func AddProduceItemToDatabase(item ProduceItem) error {
	var err error // init error for goroutine to set
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		c := make(chan error)
		go validateUUID(item.ProduceCode,c)

		if e := <- c ; e == nil {
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
		c :=  make(chan error)
		go validateUUID(produceCode, c)
		if e := <- c; e != nil {
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
*/
func validateUUID(produceCode string, err chan error)  {
	var e error
	for _, element := range database {
		if strings.ToUpper(element.ProduceCode) == strings.ToUpper(produceCode) {
			e = errors.New("Produce with this Code already exists")
		}else{
			//fmt.Println("Else Case")
			//fmt.Println(element.ProduceCode,":",produceCode)
			e = nil
		}
	}
	err <- e
	close(err)
}