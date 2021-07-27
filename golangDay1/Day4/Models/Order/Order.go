package Order

import (
	"fmt"
	"retailer-api/Config"
	Customer2 "retailer-api/Models/Customer"
	"retailer-api/Models/Product"
	"strconv"
	"sync"
	"time"

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
	order.Id=  strconv.FormatInt(time.Now().Unix(),10) // number of nanoseconds since January 1, 1970 UTC
	order.Status="In Waiting"
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	Config.MainChannel <-Config.Job{order.Id,0}
	return nil
}

func UpdateOrder(order *Order) (err error) {
	//fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

//it executes the order and update the product quantity
func ExecuteOrder(job Config.Job){
	id:=job.Id
	var order Order
	err:= GetOrderByID(&order,id)
	if err!=nil{
		fmt.Println("Failed due to order not found")
		return;
	}
	productId :=order.ProductId
	var product Product.Product
	err = Product.GetProductByID(&product, productId)
	if err != nil {
		job.Attempt=job.Attempt+1
		if job.Attempt>5{
			order.Status="Failed"
			order.Discription="Product is not available"
		} else{
			order.Status="Retrying"
			order.Status="Product is not available"
		}
		UpdateOrder(&order)
		Config.MainChannel<-job
	} else {
		customerId :=order.CustomerId
		lastExecutedTime := Customer2.GetLastExecutedTime(customerId)
		currentTime :=time.Now().Unix()
		if lastExecutedTime <0{
			job.Attempt=job.Attempt+1
			if job.Attempt>5{
				order.Status="Failed"
				order.Discription="Due to Internal Error"
				return
			} else{
				order.Status="Retrying"
				Config.MainChannel<-job
			}
			UpdateOrder(&order)
		}
		//cool down period b/w two orders of a customer
		if currentTime-lastExecutedTime <Config.CooldownPeriod{
			order.Status="Waiting"
			order.Discription="Waiting for cooldown period"
			Config.MainChannel <-job
			return
		} else{
			mutex:= &sync.Mutex{}
			mutex.Lock()
			if order.Quantity > product.Quantity {
				job.Attempt=job.Attempt+1
				if job.Attempt>5{
					order.Status="Failed"
					order.Discription="Enough Quantity of Product is not Available"
					return
				} else{
					order.Status="Retrying"
					order.Discription="Enough Quantity of Product is not Available"
					Config.MainChannel <-job
				}
			} else{
				product.Quantity=product.Quantity-order.Quantity
				Product.UpdateProduct(&product)
				order.Status="accepted"
				order.Discription="Order has been Completed"
				Customer2.UpdateLastExecutedTime(customerId,time.Now().Unix())
			}
			UpdateOrder(&order)
			mutex.Unlock()
		}
	}
}