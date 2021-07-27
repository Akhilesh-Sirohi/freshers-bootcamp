package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"retailer-api/Config"
	"retailer-api/Models/Customer"
	"retailer-api/Models/Order"
	Product2 "retailer-api/Models/Product"
	"retailer-api/Routes"
	"retailer-api/Setup"
)

var err error

func main() {

	//channel :=make(chan string,1000)
	//Config.Initiate()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	Setup.Initiate()
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Product2.Product{}, &Order.Order{}, &Customer.Customer{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
