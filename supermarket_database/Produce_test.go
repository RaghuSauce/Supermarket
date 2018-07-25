package supermarket_database

import (
	"testing"
)

//struct to hold input and output to iterate over for isEqual

type produce_isEqual_test struct {
	a     ProduceItem
	b     ProduceItem
	equal bool
}
//Struct containing the test cases for isEqual
var produce_isEqual_tests = []produce_isEqual_test{
	produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		true, //exact same struts
	},
	produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4m",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		true, //different cases for ProduceCode
	},
	produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "lettuce",
			UnitPrice:   "3.46",
		},
		true, //different cases for name
	}, produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "lettuce",
			UnitPrice:   "3.46",
		},
		true, //different cases for produce codes and Produce names
	},
	produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		false, //different produce Code
	},
	produce_isEqual_test{
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4M",
			Name:        "Apple",
			UnitPrice:   "3.46",
		},
		false, //different produce codes
	},
}

//Test for isEqual, checks to see if expected value is equal to the input values
func TestIsEqual(t *testing.T) {
	for _, elements := range produce_isEqual_tests {
		if IsEqual(elements.a, elements.b) != elements.equal {
			t.Error("something went wrong")
		}

	}
}


//struct for holding areEqual test cases
type produce_areEqual_test struct {
	a     []ProduceItem
	b     []ProduceItem
	equal bool
}


//Struct with 2 different produce items
var a_AreEqualTest_Struct = []ProduceItem{
	ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
	ProduceItem{
		ProduceCode: "E5T6-9UI3-TH15-QR88",
		Name:        "Peach",
		UnitPrice:   "2.99",
	},
}
//Struct with 2 of the same produce items
var b_AreEqualTest_Struct = []ProduceItem{
	ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
	ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
}
//Struct with 3 produce items
var c_AreEqualTest_Struct = []ProduceItem{
	ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
	ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	}, ProduceItem{
		ProduceCode: "A12T-4GH7-QPL9-3N4N",
		Name:        "Lettuce",
		UnitPrice:   "3.46",
	},
}

//Contains the test cases for the areEqual test
var produce_areEqual_tests = []produce_areEqual_test{
	produce_areEqual_test{
		a_AreEqualTest_Struct,
		b_AreEqualTest_Struct,
		false,
	},
	produce_areEqual_test{
		a_AreEqualTest_Struct,
		a_AreEqualTest_Struct,
		true,
	},
	produce_areEqual_test{
		b_AreEqualTest_Struct,
		b_AreEqualTest_Struct,
		true,
	},
	produce_areEqual_test{
		a_AreEqualTest_Struct,
		c_AreEqualTest_Struct,
		false,
	},
	produce_areEqual_test{
		b_AreEqualTest_Struct,
		c_AreEqualTest_Struct,
		false,
	},
}

//Test for AreEqual
func TestAreEqual(t *testing.T) {
	for _, elements := range produce_areEqual_tests {
		if AreEqual(elements.a, elements.b) != elements.equal {
			t.Errorf("Test for AreEqual Failed \nExpected: %t \nGot: %t",elements.equal,AreEqual(elements.a,elements.b))

		}
	}
}

func TestValidateProduceItem(t *testing.T) {
//	TODO
}
