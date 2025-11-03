package routes

import (
	"article-api/controllers"
	"article-api/middlewares"
	"article-api/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/article/", controllers.GetAllArticles)
	r.GET("/article/:id", controllers.GetArticleByID)
	r.DELETE("/article/:id", controllers.DeleteArticle)
	r.POST("/article/",
		middlewares.ValidateJSON(&models.Post{}),
		controllers.CreateArticle,
	)
	r.PUT("/article/:id",
		middlewares.ValidateJSON(&models.Post{}),
		controllers.UpdateArticle,
	)

	return r
}
