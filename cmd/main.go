package main

import (
	"api-produtos-go/controller"
	"api-produtos-go/db"
	"api-produtos-go/repository"
	"api-produtos-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnetion, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnetion)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:product_id", productController.GetProductById)
	server.DELETE("/product/:product_id", productController.DeleteProduct)

	server.Run(":8000")
}
