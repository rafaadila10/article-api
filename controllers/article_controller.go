package controllers

import (
	"article-api/database"
	"article-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	var article models.Post
	result := database.DB.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    article.Title,
		"content":  article.Content,
		"category": article.Category,
		"status":   article.Status,
	})
}

func CreateArticle(c *gin.Context) {
	val, _ := c.Get("validated_body")
	article := val.(*models.Post)

	// makesure ID zero for auto-increment
	article.ID = 0

	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post have been added.",
		"data":    article,
	})
}

// GET Posts with pagination
func GetAllArticles(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	// covert to int
	limit, err1 := strconv.Atoi(limitParam)
	offset, err2 := strconv.Atoi(offsetParam)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit and Offset must be numeric"})
		return
	}

	var posts []models.Post

	// Get data from db with limit and offset
	result := database.DB.Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Field response
	type PostResponse struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
		Status   string `json:"status"`
	}

	var response []PostResponse
	for _, post := range posts {
		response = append(response, PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			Category: post.Category,
			Status:   post.Status,
		})
	}

	c.JSON(http.StatusOK, response)
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Post
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	val, _ := c.Get("validated_body")
	input := val.(*models.Post)

	article.Title = input.Title
	article.Content = input.Content
	article.Category = input.Category
	article.Status = input.Status

	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article updated successfully.",
		"data":    article,
	})
}

// delete post
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Post

	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article deleted successfully",
	})
}
