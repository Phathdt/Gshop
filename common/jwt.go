package common

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gshop/module/users/usermodel"
)

func GenerateJWT(user *usermodel.User) (string, string, error) {
	secret := RandStringBytes(10)

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(viper.GetDuration("TOKEN_TTL") * time.Second).Unix(),
		"secret":   secret,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(viper.GetString("SIGNING_KEY")))

	return secret, signedString, err
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
