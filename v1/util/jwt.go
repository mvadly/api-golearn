package util

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ResponseTokenCreated struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthCustomClaims struct {
	Data ResponseTokenCreated
	jwt.StandardClaims
}

var signKey = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
var appName = os.Getenv("APPLICATION_NAME")
var signMethod = jwt.SigningMethodHS256

func GenerateToken(data ResponseTokenCreated) (string, error) {

	claims := &AuthCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 30).Unix(),
			Issuer:    appName,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(signMethod, claims)

	//encoded string
	t, err := token.SignedString([]byte(signKey))
	if err != nil {
		panic(err)
	}
	return t, err
}

func ValidateTokenString(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])

		}
		return []byte(signKey), nil
	})

}

func DecodeToken(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	decode := token.Claims.(jwt.MapClaims)
	data := decode["Data"].(map[string]interface{})

	return data, err
}
