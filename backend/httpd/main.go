package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"timecalcy/database"
	"timecalcy/httpd/handler"
)

func main() {
	r := gin.Default()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// init db connection
	db := database.Init()
	println("Connected to database", db)

	r.GET("/product", handler.ProductGet())
	r.POST("/workentry/create", handler.WorkentryPost())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
