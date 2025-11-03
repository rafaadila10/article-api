package main

import (
	"article-api/database"
	"article-api/routes"
	"log"
)

func main() {
	database.ConnectDatabase()
	r := routes.SetupRoutes()

	// Running server
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
