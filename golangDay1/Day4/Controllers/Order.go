package Controllers

import (
	"fmt"
	"net/http"
	"retailer-api/Models/Order"

	"github.com/gin-gonic/gin"
)


//var mainChannel chan string //will save order id of orders
//
//func Initiate(channel chan string){
//	mainChannel=channel
//	InitiateOrder()
//}
//Get All the Orders
func GetOrders(c *gin.Context) {
	var order []Order.Order
	err := Order.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//Get only one order by its id
func GetOrderByID(c *gin.Context){
	var order Order.Order
	id:=c.Params.ByName("id")
	err:= Order.GetOrderByID(&order,id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//creates the order


func CreateOrder(c *gin.Context) {
	var order Order.Order
	c.BindJSON(&order)
	//order.Id=  strconv.FormatInt(time.Now().Unix(),10) // number of nanoseconds since January 1, 1970 UTC
	//order.Status="In Waiting"
	err:= Order.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
		fmt.Println("order has been created")
		fmt.Println(order)
	}
}
//
//func InitiateOrder(){
//	go func(){
//		for{
//			select {
//				case order_id:=<-mainChannel :
//					ExecuteOrder(order_id)
//			}
//		}
//	}()
//}
//it executes the order and update the product quantity
//func ExecuteOrder(id string){
//	fmt.Println("hello ", id)
//	var order Models.Order
//	err:=Models.GetOrderByID(&order,id)
//	if err!=nil{
//		fmt.Println("order not found")
//		//for i:=0;i<5;i++{
//		//	ExecuteOrder(id);
//		//}
//		//x.AbortWithStatus(http.StatusNotFound)
//		return;
//	}
//	fmt.Println("order:")
//	fmt.Println(order)
//	product_id:=order.ProductId
//	var product Models.Product
//	err = Models.GetProductByID(&product, product_id)
//	if err != nil {
//		//x.AbortWithStatus(http.StatusNotFound)
//	} else {
//		customer_id:=order.CustomerId
//		last_executed_time:=GetLastExecutedTime(customer_id)
//		current_time:=time.Now().Unix()
//		if last_executed_time<0{
//			//x.AbortWithStatus(http.StatusNotFound)
//			order.Status="rejected due to internal error"
//			Models.UpdateOrder(&order)
//		}
//		//10 second difference b/w two orders of a customer
//		if current_time-last_executed_time <10{
//			mainChannel<-id
//			return
//		} else{
//			mutex:= &sync.Mutex{}
//			mutex.Lock()
//			if order.Quantity > product.Quantity {
//				order.Status="rejected due to Insufficient quantity in stock"
//				fmt.Println("Available Quantity of the product is less than quantity orderd")
//				//x.AbortWithStatusJSON(200, gin.H{"status": "not accepted",
//				//	"message": "Available Quantity of the product is less than quantity orderd"})
//			} else{
//				product.Quantity=product.Quantity-order.Quantity
//				Models.UpdateProduct(&product)
//				order.Status="accepted"
//				UpdateLastExecutedTime(customer_id,time.Now().Unix())
//				//x.JSON(http.StatusOK, order)
//			}
//			Models.UpdateOrder(&order)
//			mutex.Unlock()
//		}
//	}
//}
