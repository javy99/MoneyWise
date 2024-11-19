package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/javy99/MoneyWise/backend/docs/auth" // Swagger docs import
	"github.com/javy99/MoneyWise/backend/internal/auth/handler"
	"github.com/javy99/MoneyWise/backend/internal/auth/repository"
	"github.com/javy99/MoneyWise/backend/internal/auth/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Setup Swagger route
	r.GET("/swagger/auth/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup the login handler
	authRepo := repository.NewAuthRepository()         // Create a repository instance
	authService := service.NewAuthService(authRepo)    // Create service with repo
	authHandler := handler.NewAuthHandler(authService) // Create handler with service

	// Define routes
	r.POST("/login", authHandler.Login)

	r.Run(":8080") // Start the server
}
