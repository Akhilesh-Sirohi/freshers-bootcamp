package Controllers

import (
	"fmt"
	"net/http"
	"retailer-api/Models/Order"

	"github.com/gin-gonic/gin"
)


/*
GetAllOrders will return all the orders
 */
func GetAllOrders(c *gin.Context) {
	var order []Order.Order
	err := Order.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

/*
GetOrderById will return the specific order by it's id
*/
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

/*
CreateOrder will create a order and Put it into ConfigMainChannel
 */
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
