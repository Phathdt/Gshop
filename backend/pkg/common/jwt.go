package common

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gshop/module/users/usermodel"
	"gshop/pkg/config"
)

func GenerateJWT(user *usermodel.User) (string, error) {
	cfg := config.Config
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Second * time.Duration(cfg.JWT.TokenTTL)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWT.SigningKey))
}

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func GetCurrentUser(c *fiber.Ctx) *usermodel.User {
	currentUser := c.Locals("currentUser").(*usermodel.User)

	return currentUser
}
