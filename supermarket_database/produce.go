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
// Also Mutates the Produce Code and Name, wasn't sure where to place the side effect :(
func ValidateProduceItem(item *ProduceItem) error {
	item.ProduceCode = strings.ToUpper(item.ProduceCode)
	item.Name = strings.ToUpper(item.Name)

	if err := validateUUID(item.ProduceCode); err != nil {
		return nil
	}
	return errors.New("Invalid Produce Item")
}
