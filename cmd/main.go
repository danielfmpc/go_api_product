package main

import (
	"go-api/controllers"
	"go-api/db"
	"go-api/repositories"
	usecases "go-api/useCases"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Camada Repo
	ProductRepository := repositories.NewProductRepository(dbConnection)
	// Camada Case
	ProductUseCase := usecases.NewProductUseCase(ProductRepository)
	// Camada Controller
	ProductController := controllers.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProducts)

	server.Run(":8000")
}
