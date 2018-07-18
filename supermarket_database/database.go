package supermarket_database

import "github.com/pkg/errors"

var database = []ProduceItem{}

//var database = map[string]ProduceItem{}

func init() {
	//Array representing the initial state of the database
	database = []ProduceItem{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   3.46,
		},
		ProduceItem{
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			Name:        "Peach",
			UnitPrice:   2.99,
		},
		ProduceItem{
			ProduceCode: "YRT6-72AS-K736-L4AR",
			Name:        "Green Pepper",
			UnitPrice:   0.79,
		},
		ProduceItem{
			ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
			Name:        "Gala Apple",
			UnitPrice:   3.59,
		},
	}
}

//TODO Implement Concurrency
//TODO implement stubs

//Returns the all of the ProduceItems values from the database as a slice
func ListProduceItems() []ProduceItem {
	return database
}

//All produce items are assumed that they will enter via the api, thus validation will occur at the api layer
func AddProduceItemToDatabase(item ProduceItem) {
	if err := ValidateUUID(item.ProduceCode); err == nil{
		database = append(database, item)
	}
}

func RemoveProduceItemToDatabase() {
}

/*Checks to see if the Produce code already exists,
if yes, then returns a message
else returns a nil error
*/
func ValidateUUID(produceCode string) error {
	for _, element := range database{
		if element.ProduceCode == produceCode{
			return errors.New("Produce with this Code already exists")
		}
	}
	return nil
}
