package Customer

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"retailer-api/Config"
)


/*
GetCustomerByID will Fetch the specific customer by using Customer Id
 */
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}


/*
CreateCustomer will add a new Customer to the database
 */
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

/*
UpdateCustomer will update the pre-existing Customer in the database
 */
func UpdateCustomer(customer *Customer) (err error) {
	//fmt.Println(customer)
	Config.DB.Save(customer)
	return nil
}


/*
GetLastExecutedTime will return the time at which last order by the specific cutomer was processed
if the customer doesn't exist in database then it will Createcustomer and return time 0
if the error occurs it will return -1
 */
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

/*
UpdateLastExecutedTime will update the time of last executed order to the customer
*/
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