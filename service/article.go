package service

import "github.com/tanimutomo/clean-architecture-api-go/domain"

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
		// TODO: return with error
	} else if article.UserID != uid {
		// TODO: return with error
	}
	return nil
}

func (service *ArticleService) PostArticle(article domain.Article) (domain.Article, error) {
	// TODO: Assign article ID
	article, err := service.ArticleRepository.Store(article)
	return article, err
}

func (service *ArticleService) GetAllArticles(uid int) (domain.Articles, error) {
	articles, err := service.ArticleRepository.FindAllByUserID(uid)
	return articles, err
}

func (service *ArticleService) GetAllTags(uid int) (domain.Tags, error) {
	tags, err := service.TagRepository.FindAllByUserID(uid)
	return tags, err
}

func (service *ArticleService) GetArticleByID(aid int) (domain.Article, error) {
	article, err := service.ArticleRepository.FindOneByID(aid)
	return article, err
}

func (service *ArticleService) AddTag(tag domain.Tag) (domain.Tag, error) {
	// TODO: Assign a new tag ID
	tag, err := service.TagRepository.Store(tag)
	return tag, err
}

func (service *ArticleService) GetTagsByArticleID(aid int) (domain.Tags, error) {
	tags, err := service.TagRepository.FindAllByArticleID(aid)
	return tags, err
}
