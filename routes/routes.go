package routes

import (
	"article-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/article/", controllers.GetAllArticles)
	r.GET("/article/:id", controllers.GetArticleByID)
	r.POST("/article/", controllers.CreateArticle)

	return r
}
