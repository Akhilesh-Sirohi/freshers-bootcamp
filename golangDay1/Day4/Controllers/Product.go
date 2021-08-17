package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"retailer-api/Models/Product"
)

//Get All the products
func GetProducts(c *gin.Context) {
	var product []Product.Product
	err := Product.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

/*
CreateProduct will add a new product to the database
 */
func CreateProduct(c *gin.Context) {
	var product Product.Product
	c.BindJSON(&product)
	err := Product.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

/*
GetProductById will return a specific product from the database by using product Id
 */
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product.Product
	err := Product.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

/*
UpdateProduct will update the existing product in the database
*/
func UpdateProduct(c *gin.Context) {
	var product Product.Product
	id := c.Params.ByName("id")
	err := Product.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Product.UpdateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}


/*
DeleteProduct will delete a specific Product from the database by using Product Id
 */
func DeleteProduct(c *gin.Context) {
	var product Product.Product
	id := c.Params.ByName("id")
	err := Product.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}


