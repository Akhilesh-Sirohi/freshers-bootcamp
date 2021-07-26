package Routes

import (
	"retailer-api/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product-api")
	{
		//controls for product
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PUT("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)

		//controlls for order
		grp1.GET("order", Controllers.GetOrders)
		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.GetOrderByID)
	}
	return r
}
