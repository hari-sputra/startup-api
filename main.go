package main

import (
	"log"
	"os"
	"startup-api/router"

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

	router.RouteAPI(db)

}

// func authMiddleware(authService auth.AuthService, userService user.UserService) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			res := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		var token string

// 		splitToken := strings.Split(authHeader, " ")
// 		if len(splitToken) == 2 {
// 			token = splitToken[1]
// 		}

// 		t, err := authService.ValidateJWTToken(token)
// 		if err != nil {
// 			res := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		payload, ok := t.Claims.(jwt.MapClaims)
// 		if !ok || !t.Valid {
// 			res := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		userId := int(payload["user_id"].(float64))

// 		user, uErr := userService.GetUserById(userId)
// 		if uErr != nil {
// 			res := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
// 			return
// 		}

// 		c.Set("Current-User", user)

// 	}
// }
