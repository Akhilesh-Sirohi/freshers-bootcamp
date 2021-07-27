package Models

//defines the customer struct
type Customer struct {
	Id      string   `json:"id"`
	TimeofLastExecutedOrder int `json:"timeof_last_executed_order"`
}

//name of table containing customer data
func (b *Customer) TableName() string {
	return "customer"
}
