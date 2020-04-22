package database

import (
	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type ArticleRepository interface {
	Store(domain.Article) (domain.Article, error)
	FindOneByID(int) (domain.Article, error)
	FindAllByUserID(int) (domain.Articles, error)
}

type ArticleRepositoryImpl struct {
	DBHandler
}

func (repo *ArticleRepositoryImpl) Store(article domain.Article) (domain.Article, error) {
	err := repo.Create(&article).Error
	return article, err
}

func (repo *ArticleRepositoryImpl) FindOneByID(id int) (domain.Article, error) {
	var article domain.Article
	err := repo.First(&article, id).Error
	return article, err
}

func (repo *ArticleRepositoryImpl) FindAllByUserID(uid int) (domain.Articles, error) {
	var articles domain.Articles
	err := repo.Find(&articles, "id = ?", uid).Error
	return articles, err
}
