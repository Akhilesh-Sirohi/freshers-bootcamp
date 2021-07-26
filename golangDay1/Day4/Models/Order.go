package Models

import (
	"retailer-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllOrders Fetch all order data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

//GetOrderByID ... Fetch only one one by Id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

//creates the order
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}


