package routes

import (
	"startup/app/controllers"
	"startup/app/middlewares"
	"startup/app/repositories"
	"startup/app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
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
	webhookService := services.NewWebhookService(transactionService, campaignService)

	userController := controllers.NewUserController(userService, authService)
	campaignController := controllers.NewCampaignController(campaignService, campaignImageService)
	transactionController := controllers.NewTransactionController(transactionService, campaignService)
	webhookController := controllers.NewWebhookController(webhookService)

	router := gin.Default()
	router.Use(cors.Default())

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

	api.POST("/midtrans/notification", webhookController.MidtransNotification)

	/** Running in local */
	router.Run()

	/** Running with ngrok
	uncomment code below when choosing run ngrok
	*/
	// ctx := context.Background()

	// listener, err := ngrok.Listen(ctx, config.HTTPEndpoint())
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("public address: %s\n", listener.Addr())

	// if err := router.RunListener(listener); err != nil {
	// 	log.Fatalln(err)
	// }
}