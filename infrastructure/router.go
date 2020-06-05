package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(
		NewDBHandler(), NewTokenHandler(),
	)
	articleController := controllers.NewArticleController(
		NewDBHandler(),
	)

	router.POST("/signup", func(c *gin.Context) {
		userController.Signup(c)
	})
	router.POST("/login", func(c *gin.Context) {
		userController.Login(c)
	})

	inUserRouter := router.Group("/users/:userID",
		func(c *gin.Context) { userController.Authenticate(c) },
	)
	{
		// Article
		inUserRouter.POST("/articles",
			func(c *gin.Context) { articleController.PostArticle(c) },
		)
		inUserRouter.GET("/articles",
			func(c *gin.Context) { articleController.GetAllArticles(c) },
		)
		inUserRouter.GET("/tags",
			func(c *gin.Context) { articleController.GetAllTags(c) },
		)

		inArticleRouter := inUserRouter.Group("/articles/:articleID",
			func(c *gin.Context) { articleController.VerifyUser(c) },
		)
		{
			inArticleRouter.GET("/",
				func(c *gin.Context) { articleController.GetArticleByID(c) },
			)
			inArticleRouter.POST("/tags",
				func(c *gin.Context) { articleController.AddTag(c) },
			)
			inArticleRouter.GET("/tags",
				func(c *gin.Context) { articleController.GetTagsByArticleID(c) },
			)
		}
	}

	Router = router
}
