package service

import "github.com/tanimutomo/clean-architecture-api-go/domain"

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindByID(int) (domain.User, error)
}

type Token interface {
	New(domain.User) (domain.Token, error)
	Verify(domain.Token) error
}

type UserService struct {
	Repository UserRepository
	Token      Token
}

func (service *UserService) Signup(user domain.User) (domain.User, error) {
	user, err := service.Repository.Store(user)
	return user, err
}

func (service *UserService) Login(loginUser domain.LoginUser) (domain.User, domain.Token, error) {
	var token domain.Token

	user, err := service.Repository.FindByID(loginUser.ID)
	if err != nil {
		return user, token, err
	}

	// Check Password
	// TODO

	// Generate a new token
	token, err = service.Token.New(user)
	return user, token, err
}

func (service *UserService) Authenticate(token domain.Token) error {
	err := service.Token.Verify(token)
	return err
}
