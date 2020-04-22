package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type TagRepository interface {
	Store(domain.Tag) (domain.Tag, error)
	FindAllByUserID(int) (domain.Tags, error)
	FindAllByArticleID(int) (domain.Tags, error)
}

type TagRepositoryImpl struct {
	DBHandler
}

func (repo *TagRepositoryImpl) Store(tag domain.Tag) (domain.Tag, error) {
	err := repo.Create(&tag).Error
	return tag, err
}

func (repo *TagRepositoryImpl) FindAllByUserID(uid int) (domain.Tags, error) {
	var tags domain.Tags
	err := repo.Find(&tags, "user_id = ?", uid).Error
	return tags, err
}

func (repo *TagRepositoryImpl) FindAllByArticleID(aid int) (domain.Tags, error) {
	var tags domain.Tags
	err := repo.Find(&tags, "article_id = ?", aid).Error
	return tags, err
}
