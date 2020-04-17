package controllers

import (
	"net/http"
	"strconv"

	"github.com/tanimutomo/clean-architecture-api-go/domain"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/database"
	"github.com/tanimutomo/clean-architecture-api-go/service"
)

type ArticleController struct {
	Service service.ArticleService
}

func NewArticleController(dbHandler, database.DBHandler) *ArticleController {
	return &ArticleController{
		Service: service.ArticleService{
			ArticleRepository: &database.ArticleRepository{
				DBHandler: dbHandler,
			},
			TagRepository: &database.TagRepository{
				DBHandler: dbHandler,
			},
		},
	}
}

func (controller *ArticleController) VerifyUser(c Context) {
	uid, _ := strconv.Atoi(c.Param("userID"))
	aid, _ := strconv.Atoi(c.Param("articleID"))
	if err := controller.Service.VerifyUser(uid, aid); err != nil {
		// TODO: c.Abort()
		return
	}
}

func (controller *ArticleController) PostArticle(c Context) {
	uid, _ := strconv.Atoi(c.Param("userID"))
	article := domain.Article{UserID: uid}
	c.Bind(&article)

	article, err := controller.Service.PostArticle(article)
	if err != nil {
		// TODO
		return
	}
	c.JSON(http.StatusOK, article)
}

func (controller *ArticleController) GetAllArticles(c Context) {
	uid, _ := strconv.Atoi(c.Param("userID"))

	articles, err := controller.Service.GetAllArticles(uid)
	if err != nil {
		// TODO
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (controller *ArticleController) GetAllTags(c Context) {
	uid, _ := strconv.Atoi(c.Param("userID"))

	tags, err := controller.Service.GetAllTags(uid)
	if err != nil {
		// TODO
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (controller *ArticleController) GetArticleByID(c Context) {
	aid, _ := strconv.Atoi(c.Param("articleID"))

	article, err := controller.Service.GetArticleByID(aid)
	if err != nil {
		// TODO
		return
	}
	c.JSON(http.StatusOK, article)
}

func (controller *ArticleController) AddTag(c Context) {
	aid, _ := strconv.Atoi(c.Param("articleID"))

	tag := domain.Tag{ArticleID: aid}
	c.Bind(&tag)
	tag, err := controller.Service.AddTags(tag)
	if err != nil {
		// TODO
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (controller *ArticleController) GetTagsByArticleID(c Context) {
	aid, _ := strconv.Atoi(c.Param("articleID"))

	tags, err := controller.Service.GetTagsByArticleID(aid)
	if err != nil {
		// TODO
		return
	}

	c.JSON(http.StatusOK, tags)
}
