package router

import (
	"startup-api/API/user"
	"startup-api/auth"
	"startup-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAPI(db *gorm.DB) {

	// repo
	userRepository := user.NewUserRepository(db)

	// service
	userService := user.NewUserService(userRepository)
	authService := auth.NewService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// routing
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckAvailableEmail)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run(":9090")
}
