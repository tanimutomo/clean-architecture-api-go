package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	if err := LoadEnv(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	userController := controllers.NewUserController(
		NewDBHandler(), NewTokenHandler()
	)
	articleController := controllers.NewArticleController(
		NewDBHandler()
	)

	router.POST("/signup", func(c *gin.Context) {
		userController.Signup(c)
	})
	router.POST("/login", func(c *gin.Context) {
		userController.Login(c)
	})

	inUserRouter := r.Group("/users/:userID",
		func(c *gin.Context) { userController.Authenticate() }
	)
	{
		// Article
		inUserRouter.POST("/articles",
			func(c *gin.Context) { articleController.PostArticle() }
		)
		inUserRouter.GET("/articles",
			func(c *gin.Context) { articleController.GetAllArticles() }
		)
		inUserRouter.GET("/tags",
			func(c *gin.Context) { articleController.GetAllTags() }
		)

		
		inArticleRouter := r.Group("/articles/:articleID", 
			func(c *gin.Context) { articleController.VerifyUser() }
		)
		{
			inArticleRouter.GET("/", 
				func(c *gin.Context) { articleController.GetArticleByID() }
			)
			inArticleRouter.POST("/tags",
				func(c *gin.Context) { articleController.AddTag() }
			)
			inArticleRouter.GET("/tags", 
				func(c *gin.Context) { articleController.GetTagsByArticleID() }
			)
		}
	}
}
