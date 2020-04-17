package infrastructure

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/token"
)

type TokenHandler struct {}

func NewTokenHandler() token.TokenHandler {
}

func (handler *TokenHandler) Generate(uid int, username string, email string) (string, error) {
	// set header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims (json contents)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// signature
	tokenString, _ := token.SignedString([]byte(os.Getenv("SASG_SECRET")))

	return tokenString
}

func (handler *TokenHandler) Verify(tokenString, string) (error) {
	_, err := jwt.Parse(tokenString string,
		func (token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("SASG_SECRET"))
			return b, nil
		}
	)
	return error
}
