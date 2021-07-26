package Models

//defines the order struct
type Order struct{
	Id string `json:"id"`
	CustomerId int `json:"customer_id"`
	ProductId string `json:"product_id"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
}

//name of table containing order data
func (b *Order) TableName() string {
	return "order"
}