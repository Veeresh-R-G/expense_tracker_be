package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(username string, uuid string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"authorized": true,
			"username":   username,
			"uuid":       uuid,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	fmt.Printf("Token Algorithm == \n %v \n", token.Header["alg"])
	tokenString, err := token.SignedString(secretKey)
	fmt.Printf("Signed String : %s\n\n", tokenString)
	if err != nil {
		fmt.Println("Failed to Generate the JWT Token")
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Problem was to check the signing method used to 🥺

		// if token.Header["alg"] != "HS256" {
		// 	return nil, fmt.Errorf("unexpected signing algorithm: %v", token.Header["alg"])
		// }
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil

	})
	if err != nil {
		fmt.Printf("\n\nToken valid error : %v \n\n", err)
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
