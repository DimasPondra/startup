package main

import (
	"log"
	"startup/app/controllers"
	"startup/app/middlewares"
	"startup/app/repositories"
	"startup/app/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService()

	userController := controllers.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/auth/register", userController.Register)
	api.POST("/auth/login", userController.Login)
	api.POST("/auth/check-email", userController.CheckEmailAvailability)

	api.GET("/users", middlewares.AuthMiddleware(authService, userService), userController.FetchUser)
	api.POST("/users/avatar", middlewares.AuthMiddleware(authService, userService), userController.UploadAvatar)

	router.Run()
}