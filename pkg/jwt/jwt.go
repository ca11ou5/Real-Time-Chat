package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type MyJWTClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID: strconv.Itoa(userID),
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(userID),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	})

	fmt.Println(os.Getenv("SECRET_KEY"))
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
