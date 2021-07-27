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

//get the product by Id
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

//update the product
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


//Delete the product
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


