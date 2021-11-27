package util

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(arrData map[string]string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type AuthCustomClaims struct {
	ArrData map[string]string
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

var signKey = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
var appName = os.Getenv("APPLICATION_NAME")

func GenerateToken(data map[string]string) (string, error) {
	claims := &AuthCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1 / 8).Unix(),
			Issuer:    appName,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(signKey), nil
	})

}
