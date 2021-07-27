package Product


//defines the product struct
type Product struct {
	Id      string   `json:"id"`
	ProductName    string `json:"product_name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
}

//name of table containing product data
func (b *Product) TableName() string {
	return "product"
}
