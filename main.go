package main

import (
	"log"
	"os"
	"startup-api/auth"
	"startup-api/handler"
	"startup-api/user"

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

	dsn := os.Getenv("GORM_DSN")
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
