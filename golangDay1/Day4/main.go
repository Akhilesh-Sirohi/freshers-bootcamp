package main

import (
	"fmt"
	"retailer-api/Config"
	"retailer-api/Models/Customer"
	"retailer-api/Models/Order"
	Product2 "retailer-api/Models/Product"
	"retailer-api/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func Initiate(){
	Config.MainChannel =make(chan Config.Job,1000)
	go func(){
		for{
			select {
			case order_id:=<-Config.MainChannel:
				Order.ExecuteOrder(order_id)
			}
		}
	}()
}
func main() {

	//channel :=make(chan string,1000)
	//Config.Initiate()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	Initiate()
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Product2.Product{}, &Order.Order{}, &Customer.Customer{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
