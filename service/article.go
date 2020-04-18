package service

import (
	"net/http"

	"github.com/tanimutomo/clean-architecture-api-go/domain"
)

type ArticleRepository interface {
	Store(domain.Article) (domain.Article, error)
	FindOneByID(int) (domain.Article, error)
	FindAllByUserID(int) (domain.Articles, error)
}

type TagRepository interface {
	Store(domain.Tag) (domain.Tag, error)
	FindAllByUserID(int) (domain.Tags, error)
	FindAllByArticleID(int) (domain.Tags, error)
}

type ArticleService struct {
	ArticleRepository ArticleRepository
	TagRepository     TagRepository
}

func (service *ArticleService) VerifyUser(uid int, aid int) error {
	if article, err := service.ArticleRepository.FindOneByID(aid); err != nil {
		return &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if article.UserID != uid {
		return &domain.ErrorWithStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Password",
		}
	}
	return nil
}

func (service *ArticleService) PostArticle(article domain.Article) (domain.Article, error) {
	article, err := service.ArticleRepository.Store(article)
	if err != nil {
		return article, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return article, nil
}

func (service *ArticleService) GetAllArticles(uid int) (domain.Articles, error) {
	articles, err := service.ArticleRepository.FindAllByUserID(uid)
	if err != nil {
		return articles, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return articles, nil
}

func (service *ArticleService) GetAllTags(uid int) (domain.Tags, error) {
	tags, err := service.TagRepository.FindAllByUserID(uid)
	if err != nil {
		return tags, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return tags, nil
}

func (service *ArticleService) GetArticleByID(aid int) (domain.Article, error) {
	article, err := service.ArticleRepository.FindOneByID(aid)
	if err != nil {
		return article, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return article, nil
}

func (service *ArticleService) AddTag(tag domain.Tag) (domain.Tag, error) {
	tag, err := service.TagRepository.Store(tag)
	if err != nil {
		return tag, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return tag, nil
}

func (service *ArticleService) GetTagsByArticleID(aid int) (domain.Tags, error) {
	tags, err := service.TagRepository.FindAllByArticleID(aid)
	if err != nil {
		return tags, &domain.ErrorWithStatus{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return tags, nil
}
