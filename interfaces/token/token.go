package token

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type Tokenizer struct {
	TokenHandler
}

func (tokenizer *Tokenizer) New(user domain.User) (domain.Token, error) {
	var tokenString domain.Token
	generatedToken, err := tokenizer.Generate(user.ID, user.Name, user.Email)
	if err != nil {
		return tokenString, err
	}
	tokenString = domain.Token(generatedToken)
	return tokenString, nil
}

func (tokenizer *Tokenizer) Verify(token domain.Token) error {
	tokenString := string(token)
	err := tokenizer.VerifyToken(tokenString)
	return err
}
