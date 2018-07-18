package supermarket_database

//struct representing one Item of produce in the supermarket
type ProduceItem struct {
	ProduceCode string  `json:"producecode"`
	Name        string  `json:"name"`
	UnitPrice   float32 `json:"unitprice"`
}
