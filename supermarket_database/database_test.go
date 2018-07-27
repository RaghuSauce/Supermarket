package supermarket_database

import (
	"testing"
)

//Copy of values loaded into database, to create a replica of the database
var initValues = []ProduceItem{
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

func TestListProduceItems(t *testing.T) {
	c := make(chan []ProduceItem)
	go ListProduceItems(c)
	db := <-c
	testDatabase := initValues

	if !AreEqual(db, testDatabase) {
		t.Error("Database Not contain the expected initial values")
	}
}

type database_addItem_test struct {
	item  ProduceItem
	equal bool
}

var database_addItem_tests = []database_addItem_test{
	{ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
		false, //repeated code should not be allowed to be entered
	},
	{ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Apple",
		UnitPrice:   "3.46",
	},
		true,
	},
}

func TestAddProduceItemToDatabase(t *testing.T) {
	for _, element := range(database_addItem_tests){

		testDB :=  append(initValues,element.item)

		c := make(chan []ProduceItem)
		AddProduceItemToDatabase(element.item)
		go ListProduceItems(c)
		db := <- c

		if(AreEqual(testDB, db) != element.equal){
			t.Error("Error adding produce item to the databse")
		}
	}
}


func TestRemoveProduceItemFromDatabase(t *testing.T) {
	//TODO
}