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

/*
CreateOrder will change orderId to current time stamp,and will add the order to the database
It will also create the Job with the current Order Id and push it to the Config.MainChannel
 */
func CreateOrder(order *Order) (err error) {
	order.Id=  strconv.FormatInt(time.Now().Unix(),10) // number of nanoseconds since January 1, 1970 UTC
	order.Status="In Waiting"
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	Config.MainChannel <-Config.Job{order.Id,0}
	return nil
}

/*
Update Order will update the pre-existing order in the database
 */
func UpdateOrder(order *Order) (err error) {
	//fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

/*
ExecuteOrder process the Job {order_id,attempt} to complete the order
It also uses mutex to stop overlapping if multiple customer are trying to by the same product
 */
func ExecuteOrder(job Config.Job,mutex *sync.Mutex){
	id:=job.Id
	var order Order
	err:= GetOrderByID(&order,id)
	if err!=nil{
		fmt.Println("Failed due to order not found")
		return
	}
	//get the product using the productId
	productId :=order.ProductId
	var product Product.Product
	err = Product.GetProductByID(&product, productId)


	if err != nil {
		//If the product isn't present it will increase attempt
		job.Attempt=job.Attempt+1

		/*If the number of attempt is more than 5 than the Job will fail
		 else the Job will be retried (pushed back to the MainChannel
		 */
		if job.Attempt>5{
			order.Status="Failed"
			order.Discription="Product is not available"
		} else{
			order.Status="Retrying"
			order.Discription="Product is not available"
		}

		//UpdateOrder is Used to update status and descreption of the order
		UpdateOrder(&order)
		Config.MainChannel<-job
	} else {
		customerId :=order.CustomerId

		/*
			GetLastExecutedTime will return the time at which last order by the specific cutomer was processed
			if the customer doesn't exist in database then it will Createcustomer and return time 0
			if the error occurs it will return -1
		*/
		lastExecutedTime := Customer2.GetLastExecutedTime(customerId)
		currentTime :=time.Now().Unix()

		/*
		    if the lastExecutedTime is less than 0 that means there was a error while executing GetLastExecutedTime
		 */
		if lastExecutedTime <0{
			//So it will increase the attempt of the job
			job.Attempt=job.Attempt+1


			/*
				If the number of attempt is more than 5 than the Job will fail
				else the Job will be retried (pushed back to the MainChannel
			*/
			if job.Attempt>5{
				order.Status="Failed"
				order.Discription="Due to Internal Error"
			} else{
				order.Status="Retrying"
				order.Discription="Due to Internal Error"
				Config.MainChannel<-job
			}
			//UpdateOrder is Used to update status and descreption of the order
			UpdateOrder(&order)
			return
		}

		/*
			cool down period is the mandatory time gap b/w two orders/Job of same customer
			so if the cool down period hasn't been completed so the Job will be pushed back to the channel to be
			executed again
		 */
		if currentTime-lastExecutedTime <Config.CooldownPeriod{
			order.Status="Waiting"
			order.Discription="Waiting for cooldown period"
			Config.MainChannel <-job
		} else{
			//mutex.Lock() to stop overlapping if multiple customer are trying to by the same product
			mutex.Lock()
			if order.Quantity > product.Quantity {
				job.Attempt=job.Attempt+1
				if job.Attempt>5{
					order.Status="Failed"
					order.Discription="Enough Quantity of Product is not Available"
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
			mutex.Unlock()
		}
		UpdateOrder(&order)
	}
}