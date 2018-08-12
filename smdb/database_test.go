package smdb

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

type database_getOne_test struct {
	testName      string //The name of the test case
	itemCode      string //The Item code we are looking for
	isValidSearch bool   //Are we expecting to find the value
}

var database_getOne_tests = []database_getOne_test{
	{
		"Valid Search 1",
		"A12T-4GH7-QPL9-3N4M",
		true,
	},
	{
		"Valid Search 2",
		"E5T6-9UI3-TH15-QR88",
		true,
	},
	{
		"Search for non existent item",
		"E5T6-9UI3-TH15-QR8N",
		false,
	},
	{
		"Search for Erroneous Item",
		"12",
		false,
	},
}

func TestGetOneProduceItem(t *testing.T) {
	for _, element := range database_getOne_tests { //For each item in the test cases
		itemExits, item := GetProduceItem(element.itemCode) //get the produce item and the bool indicating if it exists or not
		if itemExits == element.isValidSearch {             //if the retrieved bool and expected bool are equal && the produce codes of both items are the same, do nothing
			if item.ProduceCode == element.itemCode {
				//	do nothing
			}
		} else { //Else Error out
			t.Errorf("Error finding Produce item\nFailed:%s"+
				"Expected ItemCode:%s\nGot:%s",
				element.testName, element.itemCode, item.ProduceCode)
		}
	}
}

func TestListProduceItems(t *testing.T) {
	c := make(chan []ProduceItem) //make a chan to the database
	go ListProduceItems(c)        // populate the channel with db items
	db := <-c                     //dump channel to var
	testDatabase := initValues    //Create test database for comparison

	if !AreEqual(db, testDatabase) { //check if they are equal
		t.Error("Database does not contain the expected initial values")
	}
}

type database_addItem_test struct {
	testName string      //The name of the test scenario
	item     ProduceItem //The Produce item that will be used in the test
	equal    bool        //The expected result of the database transaction
}

var database_addItem_tests = []database_addItem_test{
	{
		"Adding new valid ProduceItem",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Apple",
			UnitPrice:   "3.46",
		},
		true,
	},

	{
		"Adding existing ProduceItem",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		false, //repeated code should not be allowed to be entered
	},
	{
		"Adding Invalid Item",
		ProduceItem{
			ProduceCode: "12",
			Name:        "*",
			UnitPrice:   "12.000",
		}, true, //Item is Invalid thus should not be added to the db
	},
}

func TestAddProduceItemToDatabase(t *testing.T) {
	for _, element := range database_addItem_tests {
		testDB := append(initValues, element.item) // Create an array that appends all test produce items to itself
		c := make(chan []ProduceItem)
		AddProduceItemToDatabase(element.item) //Attempt to add the produce item to the the database, run it normally to make sure we attempt to add the item
		go ListProduceItems(c)
		db := <-c                                  //Get the current database
		if AreEqual(testDB, db) != element.equal { //Check to see if the databases should be equal according to the the test logic
			t.Errorf("Error adding produce item to the databse \n Failed Test Case:%s", element.testName)

		}
		testDB = nil //Reset the the test database for the next test case
		ResetDB()    //Reset the test database for the next test case
	}
}

type database_removeItem_test struct {
	testName       string        //The name of the test scenario
	item           ProduceItem   //The Produce item that will be used in the test
	expectedOutput []ProduceItem // the list of items we are expecting back as a a result
	expectedResult bool          //The expected result of the database transaction
}

var database_RemoveItems_test = []database_removeItem_test{
	{
		"Invalid Removal: Tried to remove non existent code",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]ProduceItem{
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
		},
		false,
	},
	{
		"Valid Removal",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]ProduceItem{
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
		},
		true,
	},
}

func TestRemoveProduceItemFromDatabase(t *testing.T) {

	for _, element := range database_RemoveItems_test {
		RemoveProduceItemFromDatabase(element.item.ProduceCode) //attempt to remove an item from the database
		c := make(chan []ProduceItem)
		go ListProduceItems(c)
		db := <-c                                                           //fetch the current database
		if AreEqual(db, element.expectedOutput) != element.expectedResult { //Check to see if the test results and the actual are supposed to be equal
			t.Errorf("Failed to remove produce item, \n Failed:%s\nExpected Result and actual result for removing element do not match\nExpected:\n%s\nGot:\n%s",
				element.testName, element.expectedOutput, db)
		}
		ResetDB() //Reset the database for future tests
	}
}
