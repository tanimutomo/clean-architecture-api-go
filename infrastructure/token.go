package infrastructure

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/token"
)

type TokenHandler struct{}

func NewTokenHandler() token.TokenHandler {
	return new(TokenHandler)
}

func (handler *TokenHandler) Generate(uid int, username string, email string) (string, error) {
	// set header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims (json contents)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["username"] = username
	claims["email"] = email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// signature
	tokenString, err := token.SignedString([]byte(os.Getenv("CAAG_SECRET")))

	return tokenString, err
}

func (handler *TokenHandler) Verify(tokenString string) error {
	_, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("CAAG_SECRET"))
			return b, nil
		},
	)
	return err
}
