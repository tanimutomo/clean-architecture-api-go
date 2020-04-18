package token

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type Token struct {
	TokenHandler
}

func (token *Token) New(user domain.User) (domain.Token, error) {
	generatedToken, err := token.Generate(user.ID, user.Name, user.Email)
	tokenString := domain.Token(generatedToken)
	return tokenString, err
}

func (token *Token) Verify(tokenString domain.Token) error {
	err := token.Verify(tokenString)
	return err
}
