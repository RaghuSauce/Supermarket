package supermarket_database

import (
	"strings"
	"errors"
)

//struct representing one Item of produce in the supermarket
type ProduceItem struct {
	ProduceCode string  `json:"producecode"`
	Name        string  `json:"name"`
	UnitPrice   float32 `json:"unitprice"`
}

//Validates the incoming produce Item,

//TODO limit precision on unit price to two decimal plalces
//TODO add validation for Produce Code
func ValidateProduceItem(item *ProduceItem) error {
	item.ProduceCode = strings.ToUpper(item.ProduceCode)
	item.Name = strings.ToUpper(item.Name)

	if err := ValidateUUID(item.ProduceCode); err != nil {
		return nil
	}
	return errors.New("Invalid Produce Item")
}
