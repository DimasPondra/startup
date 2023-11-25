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
	campaignRepo := repositories.NewCampaignRepository(db)
	campaignImageRepo := repositories.NewCampaignImageRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService()
	campaignService := services.NewCampaignSevice(campaignRepo)
	campaignImageService := services.NewCampaignImageService(campaignImageRepo)
	paymentService := services.NewPaymentService()
	transactionService := services.NewTransactionService(transactionRepo, paymentService)

	userController := controllers.NewUserController(userService, authService)
	campaignController := controllers.NewCampaignController(campaignService, campaignImageService)
	transactionController := controllers.NewTransactionController(transactionService, campaignService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/auth/register", userController.Register)
	api.POST("/auth/login", userController.Login)
	api.POST("/auth/check-email", userController.CheckEmailAvailability)

	api.GET("/users", middlewares.AuthMiddleware(authService, userService), userController.FetchUser)
	api.POST("/users/avatar", middlewares.AuthMiddleware(authService, userService), userController.UploadAvatar)

	api.GET("/campaigns", campaignController.Index)
	api.GET("/campaigns/:slug/show", campaignController.Show)
	api.POST("/campaigns/store", middlewares.AuthMiddleware(authService, userService), campaignController.Store)
	api.PATCH("/campaigns/:slug/update", middlewares.AuthMiddleware(authService, userService), campaignController.Update)
	api.POST("/campaigns/:slug/upload-images", middlewares.AuthMiddleware(authService, userService), campaignController.UploadImages)

	api.GET("/transactions", middlewares.AuthMiddleware(authService, userService), transactionController.Index)
	api.POST("/transactions/store", middlewares.AuthMiddleware(authService, userService), transactionController.Store)

	router.Run()
}