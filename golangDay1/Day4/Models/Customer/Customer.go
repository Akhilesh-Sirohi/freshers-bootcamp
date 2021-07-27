package Customer

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"retailer-api/Config"
)


//GetCustomerByID ... Fetch only one one by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
//creates the customer
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//update the customer
func UpdateCustomer(customer *Customer) (err error) {
	//fmt.Println(customer)
	Config.DB.Save(customer)
	return nil
}

func GetLastExecutedTime(id string) int64{
	var customer Customer
	err:= GetCustomerByID(&customer,id)
	if err!=nil{
		customer.Id=id
		err= CreateCustomer(&customer)
		if err!=nil{
			fmt.Println(err.Error())
			return -1;
		}
		return 0;
	}
	return customer.TimeofLastExecutedOrder
}

func UpdateLastExecutedTime(id string,time int64){
	var customer Customer
	err:= GetCustomerByID(&customer,id)
	if err!=nil{
		fmt.Println(err.Error())
	} else{
		customer.TimeofLastExecutedOrder=time
		err= UpdateCustomer(&customer)
		if err!=nil{
			fmt.Println(err.Error())
		}
	}
}