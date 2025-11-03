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
	r.PUT("/article/:id", controllers.UpdateArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)

	return r
}
