package service

import (
	"net/http"

	"github.com/tanimutomo/clean-architecture-api-go/domain"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/database"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/token"
)

type UserService struct {
	Repository database.UserRepository
	Tokenizer  token.Tokenizer
}

func (service *UserService) Signup(user domain.User) (domain.User, error) {
	user, err := service.Repository.Store(user)
	if err != nil {
		return user, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return user, nil
}

func (service *UserService) Login(loginUser domain.LoginUser) (domain.User, domain.Token, error) {
	var token domain.Token

	user, err := service.Repository.FindByID(loginUser.ID)
	if err != nil {
		return user, token, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if user.Password != loginUser.Password {
		return user, token, &domain.ErrorWithStatus{
			Status:  http.StatusBadRequest,
			Message: "Invalid Password",
		}
	}

	// Generate a new token
	token, err = service.Tokenizer.New(user)
	if err != nil {
		return user, token, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return user, token, nil
}

func (service *UserService) Authenticate(token domain.Token) error {
	if err := service.Tokenizer.Verify(token); err != nil {
		return &domain.ErrorWithStatus{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}
