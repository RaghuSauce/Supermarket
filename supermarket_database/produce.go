package supermarket_database

import (
	"regexp"
)

//struct representing one Item of produce in the supermarket
type ProduceItem struct {
	ProduceCode string `json:"producecode"`
	Name        string `json:"name"`
	UnitPrice   string `json:"unitprice"`
}

//Regex for Allowed Values for ProduceItems
//TODO Implement returning proper status codes
const (
	PRODUCECODEREGEX = "^(([A-Z 0-9 a-z]){4}-){3}(([A-Z 0-9 a-z]){4})$"
	NAMEREGEX        = "^(([A-Za-z0-9])*)(\\s[A-Za-z0-9 ]*)?$"
	UNITPRICEREGEX   = "^[0-9]+(\\.[0-9]{1,2})?$"
)

//Validates the incoming produce Item,
func ValidateProduceItem(item ProduceItem) (bool, error) {

	var isValidProduceItem bool                          //bool to represent if the produce Item is valid
	r, err := regexp.Compile(PRODUCECODEREGEX)           // compile The Produce Code Regex
	isValidProduceItem = r.MatchString(item.ProduceCode) //determine if produce code is valid

	r, err = regexp.Compile(NAMEREGEX)
	isValidProduceItem = isValidProduceItem && r.MatchString(item.Name) //determine if the produce code and name are valid

	r, err = regexp.Compile(UNITPRICEREGEX)
	//unitPriceString := strconv.FormatFloat(item.UnitPrice, 'f', 32, 64)       //convert the float to string with enough precision
	//fmt.Println(unitPriceString)
	isValidProduceItem = isValidProduceItem && r.MatchString(item.UnitPrice) //determine if the price, name and produce code are valid

	return isValidProduceItem, err
}
