package supermarket_database

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

//struct representing one Item of produce in the supermarket
type ProduceItem struct {
	ProduceCode string `json:"producecode"`
	Name        string `json:"name"`
	UnitPrice   string `json:"unitprice"`
}

//Regex for Allowed Values for ProduceItems
const (
	PRODUCECODEREGEX = "^(([A-Z 0-9 a-z]){4}-){3}(([A-Z 0-9 a-z]){4})$"
	NAMEREGEX        = "^(([A-Za-z0-9])*)(\\s[A-Za-z0-9 ]*)?$"
	UNITPRICEREGEX   = "^[0-9]+(\\.[0-9]{1,2})?$"
)

func AreEqual(a []ProduceItem, b []ProduceItem) bool {
	// if the slices are the same size
	areEqual := false
	if len(a) == len(b) {
		areEqual = true
		for i, _ := range a {
			areEqual = areEqual && IsEqual(a[i], b[i])
		}
	}
	return areEqual
}
func IsEqual(a ProduceItem, b ProduceItem) bool {
	return (strings.ToUpper(a.ProduceCode) == strings.ToUpper(b.ProduceCode)) &&
		(strings.ToUpper(a.Name) == strings.ToUpper(b.Name)) &&
		(a.UnitPrice == b.UnitPrice)
}

var (
	INVALID_PRODUCE_CODE_ERROR error = errors.New("Invalid Produce Code")
	INVALID_PRODUCE_NAME_ERROR error = errors.New("Invalid Produce Name")
	INVALID_PRICE_ERROR        error = errors.New("Invalid Price")
)

//Validates the incoming produce Item,
func ValidateProduceItem(item ProduceItem) (bool, []error) {

	var isValidProduceItem bool //bool to represent if the produce Item is valid
	//var errString string
	r, _ := regexp.Compile(PRODUCECODEREGEX)              // compile The Produce Code Regex
	isValidProduceCode := r.MatchString(item.ProduceCode) //determine if produce code is valid

	r, _ = regexp.Compile(NAMEREGEX)
	isValidName := r.MatchString(item.Name) //determine if the produce code and name are valid

	r, _ = regexp.Compile(UNITPRICEREGEX)
	//unitPriceString := strconv.FormatFloat(item.UnitPrice, 'f', 32, 64)       //convert the float to string with enough precision
	//fmt.Println(unitPriceString)
	isValidPrice := r.MatchString(item.UnitPrice) //determine if the price, name and produce code are valid

	isValidProduceItem = isValidName && isValidPrice && isValidProduceCode

	var errs []error
	if !isValidProduceItem {
		if !isValidProduceCode {
			errs = append(errs, INVALID_PRODUCE_CODE_ERROR)
		}
		if !isValidName {
			errs = append(errs, INVALID_PRODUCE_NAME_ERROR)
		}
		if !isValidPrice {
			errs = append(errs, INVALID_PRICE_ERROR)
		}
	}
	return isValidProduceItem, errs
}
