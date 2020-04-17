package token

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type Token struct {
	TokenHandler
}

func (token *Token) New(user domain.User) (domain.Token, error) {
	var generatedToken &domain.Token
	tokenString, err := token.Generate(user.ID, user.Name, user.Email)
	generatedToken = *tokenString
	return generatedToken, err
}

func (token *Token) Verify(tokenString domain.Token) error {
	err := token.Verify(tokenString)
	return err
}
