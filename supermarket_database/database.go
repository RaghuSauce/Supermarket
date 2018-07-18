package supermarket_database

//var database = []ProduceItem{}

var database = map[string]ProduceItem{}

func init() {
	//Array representing the initial state of the database
	var initDatabase = []ProduceItem{
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

	database = make(map[string]ProduceItem) //init the database
	for _, element := range initDatabase {  //add all the initial structs to the database
		produceCode := element.ProduceCode
		database[produceCode] = element

	}
}

//TODO Implement Concurrency
//TODO implement stubs

//Returns the all of the ProduceItems values from the database as a slice
func ListProduceItems() []ProduceItem {
	var databaseList = []ProduceItem{}
	for _, element := range database {
		databaseList = append(databaseList, element)
	}
	return databaseList
}

func AddProduceItem(item ProduceItem) {
}

func RemoveProduceItem() {
}


//TODO actually check the produceCode
func validateUUID(id string) error {
	return nil
}
