package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthService interface {
	GenerateJWTToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateJWTToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
