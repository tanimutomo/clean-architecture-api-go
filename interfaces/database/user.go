package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type UserRepository struct {
	DBHandler
}

func (repo *UserRepository) FindById(id int) (domain.User, error) {
	var user domain.User
	err := repo.First(&user, id).Error
	return user, err
}

func (repo *UserRepository) Store(user domain.User) (domain.User, error) {
	err := repo.Create(&user).Error
	return user, err
}
