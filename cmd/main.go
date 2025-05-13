package main

import (
	"go-api/auth"
	"go-api/controller"
	"go-api/db"
	"go-api/env"
	"go-api/middleware"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	env.Load()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	TokenMiddleware := middleware.NewTokenMiddleware(env.JwtSecret.GetValue())
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/login", func(ctx *gin.Context) {
		auth.Login(ctx, env.JwtSecret.GetValue())
	})

	server.Use(TokenMiddleware.Auth())

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
