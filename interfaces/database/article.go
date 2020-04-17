package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type ArticleRepository struct {
	DBHandler
}

func (repo *ArticleRepository) Store(article domain.Article) (domain.Article, error) {
	err := repo.Create(&article).Error
	return article, err
}

func (repo *ArticleRepository) FindOneById(id int) (domain.Article, error) {
	var article domain.Article
	err := repo.First(&article, id).Error
	return article, err
}

func (repo *ArticleRepository) FindAllByUserID(uid int) (domain.Articles, error) {
	var articles domain.Articles
	err := repo.Find(&articles, "id = ?", uid).Error
	return articles, err
}
