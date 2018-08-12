package smdb

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
			t.Errorf("Test ProduceCode Equality Failed \n Comparing (%s | %s)", elements.a, elements.b)
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
			t.Errorf("Test for AreEqual Failed \nExpected: %t \nGot: %t", elements.equal, AreEqual(elements.a, elements.b))

		}
	}
}

var emptyErrorSet = make([]error, 0)

//struct to hold test scenarios for ValidateProduceItems
type produce_ValidateProduceItem_struct struct {
	TestName    string      //The name of the test case
	produceItem ProduceItem // the item being validated
	errs        []error     //the error expected
	isValid     bool        //the expected result from the test

}

//tests scenarios for ValidateProduceItem func
var produce_TestValidateProduceItem_tests = []produce_ValidateProduceItem_struct{

	produce_ValidateProduceItem_struct{
		"Valid Produce Item",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		nil,
		true,
	},

	produce_ValidateProduceItem_struct{
		"Produce is lowercase",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Produce Is Mixed Case",
		ProduceItem{
			ProduceCode: "a12t-4Gh7-Ppl9-3n4n",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},

	produce_ValidateProduceItem_struct{
		"Produce is lowercase",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},

	produce_ValidateProduceItem_struct{
		"Produce Code is missing digit the the first group",
		ProduceItem{
			ProduceCode: "A12-4GH7-QPL9-3N4M",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		}, []error{
			INVALID_PRODUCE_CODE_ERROR,
		},
		false,
	},

	produce_ValidateProduceItem_struct{
		"Produce Item is missing digit from the 2nd group",
		ProduceItem{
			ProduceCode: "A12T-4GH-QPL9-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Produce Item is missing digit from the 3nd group",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL-3N4N",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Produce Item is missing digit from the 4th group",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL9-3N4",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"ProduceItems does not fit the expected format",
		ProduceItem{
			ProduceCode: "A12T-4GH7-QPL993N45",
			Name:        "Lettuce",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"C_ (())",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "C_ (())",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_NAME_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Empty Name",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Space Name",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        " ",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Lowercase Name",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "lettuce",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"UpperCase name",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4n",
			Name:        "uppercase name",
			UnitPrice:   "3.46",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Produce Code and Invalid Name",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4",
			Name:        "(👚)",
			UnitPrice:   "3.46",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
			INVALID_PRODUCE_NAME_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Valid Price without cents",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "300",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Valid Price with Precision 1 on cents",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "300.0",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Valid Price with Precision 2 on cents",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "300.00",
		},
		[]error{},
		true,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Price too many decimal points",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "300.000",
		},
		[]error{
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Empty Price",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "",
		},
		[]error{
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Negative Price",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "-20.00",
		},
		[]error{
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Produce Code and Invalid UnitPrice",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4",
			Name:        "Shirt",
			UnitPrice:   "300.000",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Name and Invalid Price",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shir-",
			UnitPrice:   "300.000",
		},
		[]error{
			INVALID_PRODUCE_NAME_ERROR,
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Price too many decimal points",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n43",
			Name:        "Shirt",
			UnitPrice:   "300.000",
		},
		[]error{
			INVALID_PRICE_ERROR,
		},
		false,
	},
	produce_ValidateProduceItem_struct{
		"Invalid Produce Code, Invalid Name, Invalid Unit Price",
		ProduceItem{
			ProduceCode: "a12t-4gh7-qpl9-3n4",
			Name:        "(👚)",
			UnitPrice:   "300.000",
		},
		[]error{
			INVALID_PRODUCE_CODE_ERROR,
			INVALID_PRODUCE_NAME_ERROR,
			INVALID_PRICE_ERROR,
		},
		false,
	},
}

//function to test validation of produce items
func TestValidateProduceItem(t *testing.T) {
	var errs []error
	var isValid bool

	for _, element := range produce_TestValidateProduceItem_tests {
		isValid, errs = ValidateProduceItem(element.produceItem)
		if !isValid == element.isValid || !checkErrorsEqual(errs, element.errs) {

			t.Errorf("Test for Validate Produce Item Failed On test: %s"+
				" \nExpected bool:%t\nGot bool:%t"+
				" \nExpected errors:%s\nGot errors:%s",
				element.TestName, element.isValid, isValid, element.errs, errs)
		}
	}

}
func checkErrorsEqual(a []error, b []error) bool {
	isEqual := false
	if len(a) == len(b) {
		isEqual = true
		for i, _ := range a {
			if a[i] != b[i] {
				isEqual = false
			}
		}
	}
	return isEqual
}
