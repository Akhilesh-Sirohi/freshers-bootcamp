package Models

import (
	"retailer-api/Config"

	_ "github.com/go-sql-driver/mysql"
)


//GetCustomerByID ... Fetch only one one by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
