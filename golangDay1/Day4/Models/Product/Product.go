package Product

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"retailer-api/Config"
)

//GetAllProducts Fetch all user data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one product by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProduct ... Update product
func UpdateProduct(product *Product) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

//DeleteProduct ... Delete Product
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}

