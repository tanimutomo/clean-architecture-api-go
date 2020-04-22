package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindByID(int) (domain.User, error)
}

type UserRepositoryImpl struct {
	DBHandler
}

func (repo *UserRepositoryImpl) FindByID(id int) (domain.User, error) {
	var user domain.User
	err := repo.First(&user, id).Error
	return user, err
}

func (repo *UserRepositoryImpl) Store(user domain.User) (domain.User, error) {
	err := repo.Create(&user).Error
	return user, err
}
