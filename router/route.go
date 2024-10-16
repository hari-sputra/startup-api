package router

import (
	"startup-api/API/campaign"
	"startup-api/API/user"
	"startup-api/auth"
	"startup-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAPI(db *gorm.DB) {

	// repo
	userRepository := user.NewUserRepository(db)
	campaignRepository := campaign.NewCampaignRepository(db)

	// service
	userService := user.NewUserService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewCampaignService(campaignRepository)

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// routing
	// user
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckAvailableEmail)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	// campaign
	api.GET("/campaigns", campaignHandler.GetCampaign)

	router.Run(":9090")
}
