package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signatureKey = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
var applicationName = os.Getenv("APPLICATION_NAME")
var signingMethod = jwt.SigningMethodHS256

type MyClaims struct {
	jwt.StandardClaims
	Data map[string]string
}

func CreateToken(arrData map[string]string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		},
		Data: arrData,
	}

	token := jwt.NewWithClaims(
		signingMethod,
		claims,
	)

	signedToken, err := token.SignedString(signatureKey)
	// tokenString, _ := json.Marshal(map[string]string{"token": signedToken})

	return string(signedToken), err
}

func ValidateToken() {

}
