package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mateochauchet/go-ecom/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationSec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  strconv.Itoa(userID),
		"expiedAt": time.Now().Add(expiration).Unix(),
	})

	stringToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return stringToken, nil
}
