package Controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
	//"github.com/retailer-api/Models"
	"retailer-api/Models"

	"github.com/gin-gonic/gin"
)


var mainChannel chan string //will save order id of orders

func Initiate(channel chan string){
	mainChannel=channel
	InitiateOrder()
}
//Get All the Orders
func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//Get only one order by its id
func GetOrderByID(c *gin.Context){
	var order Models.Order
	id:=c.Params.ByName("id")
	err:=Models.GetOrderByID(&order,id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//creates the order

var x *gin.Context
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	order.Id=  strconv.FormatInt(time.Now().Unix(),10) // number of nanoseconds since January 1, 1970 UTC
	order.Status="In Waiting"
	err:= Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
		x=c
		fmt.Println("order has been created")
		fmt.Println(order)
		mainChannel<-order.Id
	}


}

func InitiateOrder(){
	go func(){
		for{
			select {
				case order_id:=<-mainChannel :
					ExecuteOrder(order_id)
			}
		}
	}()
}
//it executes the order and update the product quantity
func ExecuteOrder(id string){
	fmt.Println("hello ", id)
	var order Models.Order
	err:=Models.GetOrderByID(&order,id)
	if err!=nil{
		fmt.Println("order not found")
		x.AbortWithStatus(http.StatusNotFound)
		return;
	}
	fmt.Println("order:")
	fmt.Println(order)
	product_id:=order.ProductId
	var product Models.Product
	err = Models.GetProductByID(&product, product_id)
	if err != nil {
		x.AbortWithStatus(http.StatusNotFound)
	} else {
		mutex:= &sync.Mutex{}
		mutex.Lock()
		if order.Quantity > product.Quantity {
			x.AbortWithStatusJSON(200, gin.H{"status": "not accepted",
				"message": "Available Quantity of the product is less than quantity orderd"})
		} else{
			product.Quantity=product.Quantity-order.Quantity
			Models.UpdateProduct(&product)
			order.Status="accepted"
			x.JSON(http.StatusOK, order)
		}
		mutex.Unlock()
	}
}
