package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type TagRepository struct {
	DBHandler
}

func (repo *TagRepository) Store(tag domain.Tag) (domain.Tag, error) {
	err := repo.Create(&tag).Error
	return tag, err
}

func (repo *TagRepository) FindAllByUserID(uid int) (domain.Tag, error) {
	var tag domain.Tag
	err := repo.Find(&tag, "user_id = ?", uid).Error
	return tag, err
}

func (repo *TagRepository) FindAllByArticleID(aid int) (domain.Tags, error) {
	var tags domain.Tags
	err := repo.Find(&tags, "article_id = ?", aid).Error
	return tags, err
}