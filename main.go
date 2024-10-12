package main

import (
	"log"
	"startup-api/auth"
	"startup-api/handler"
	"startup-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/startup_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// user
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	// auth
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckAvailableEmail)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run(":9090")

}
