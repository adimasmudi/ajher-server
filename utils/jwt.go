package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(userID int) (interface{}, error) {
	accessToken, err := generateAccessToken(userID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken(userID)

	if err != nil {
		return nil, err
	}

	return gin.H{"accessToken": accessToken, "refreshToken": refreshToken}, err
}

func generateAccessToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func generateRefreshToken(userID int) (string, error) {
	claim := jwt.MapClaims{}

	claim["exp"] = time.Now().Add(time.Hour * 2160).Unix() // refresh token active for 3 months

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
