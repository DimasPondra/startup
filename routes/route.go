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
	roleRepo := repositories.NewRoleRepository(db)

	roleService := services.NewRoleService(roleRepo)
	userService := services.NewUserService(userRepo, roleService)
	authService := services.NewAuthService()
	campaignService := services.NewCampaignSevice(campaignRepo)
	campaignImageService := services.NewCampaignImageService(campaignImageRepo)
	paymentService := services.NewPaymentService()
	transactionService := services.NewTransactionService(transactionRepo, paymentService)
	webhookService := services.NewWebhookService(transactionService, campaignService)

	userController := controllers.NewUserController(userService, authService)
	campaignController := controllers.NewCampaignController(campaignService, campaignImageService, transactionService)
	transactionController := controllers.NewTransactionController(transactionService, campaignService)
	webhookController := controllers.NewWebhookController(webhookService)
	roleController := controllers.NewRoleController(roleService)

	authMiddleware := middlewares.AuthMiddleware(authService, userService)

	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/auth/register", userController.Register)
	api.POST("/auth/login", userController.Login)
	api.POST("/auth/check-email", userController.CheckEmailAvailability)

	api.GET("/users", authMiddleware, userController.FetchUser)
	api.POST("/users/avatar", authMiddleware, userController.UploadAvatar)

	api.GET("/campaigns", campaignController.Index)
	api.GET("/campaigns/:slug/show", campaignController.Show)
	api.POST("/campaigns/store", authMiddleware, campaignController.Store)
	api.PATCH("/campaigns/:slug/update", authMiddleware, campaignController.Update)
	api.POST("/campaigns/:slug/upload-images", authMiddleware, campaignController.UploadImages)
	api.GET("/campaigns/:slug/transactions", authMiddleware, campaignController.ShowTransactions)

	api.GET("/transactions", authMiddleware, transactionController.Index)
	api.POST("/transactions/store", authMiddleware, transactionController.Store)

	api.POST("/midtrans/notification", webhookController.MidtransNotification)

	api.GET("/roles", roleController.Index)
	api.POST("/roles/store", roleController.Store)

	// Running in local
	router.Run()

	/*
		Running with ngrok
		uncomment code below when choosing run ngrok
	*/
	/*
		ctx := context.Background()

		listener, err := ngrok.Listen(ctx, config.HTTPEndpoint())
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("public address: %s\n", listener.Addr())

		if err := router.RunListener(listener); err != nil {
			log.Fatalln(err)
		}
	*/
}
