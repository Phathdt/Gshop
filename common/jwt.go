package common

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
	"gshop/module/users/usrmodel"
)

func GenerateJWT(user *usrmodel.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(viper.GetDuration("TOKEN_TTL") * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("SIGNING_KEY")))
}
