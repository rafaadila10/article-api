package main

import (
	"article-api/database"
	"article-api/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173"}, // local only
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 🧩 Setup routes
	routes.SetupRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
