package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

var products []Product

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/products", func(c *gin.Context) {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
		c.HTML(http.StatusOK, "products.html", gin.H{"products": products})
	})

	router.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
		c.JSON(http.StatusCreated, gin.H{"status": "Product created successfully"})
	})

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.Run(":5500")
}
