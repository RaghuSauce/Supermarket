package supermarket_database

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

func AddProduceItem(item ProduceItem) {

}

func RemoveProduceItem() {
}

/*Checks to see if the Produce code actually exists,
if yes then return a nil error
else return a message
*/
func ValidateUUID(id string) error {
	return nil
}
