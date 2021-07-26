package Controllers

import (
	"fmt"
	"net/http"
	//"github.com/retailer-api/Models"
	"retailer-api/Models"

	"github.com/gin-gonic/gin"
)

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
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	order.Status="In Waiting"
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		ExecuteOrder(&order,c)
	}
}

//it executes the order and update the product quantity
func ExecuteOrder(order *Models.Order, c *gin.Context){
	product_id:=order.ProductId
	var product Models.Product
	err := Models.GetProductByID(&product, product_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if order.Quantity > product.Quantity {
		c.AbortWithStatusJSON(200, gin.H{"status": "not accepted",
			"message": "Available Quantity of the product is less than quantity orderd"})

	} else{
		product.Quantity=product.Quantity-order.Quantity
		Models.UpdateProduct(&product)
		order.Status="accepted"
		c.JSON(http.StatusOK, order)
	}
}

