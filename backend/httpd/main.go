package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"timecalcy/database"
	"timecalcy/httpd/controller"
	"timecalcy/httpd/handler"
)

func main() {
	r := gin.Default()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	db := database.Init()
	println("Connected to database", db)

	workEntryCtrl := new(controller.WorkEntry)

	r.GET("/product", handler.ProductGet())
	r.POST("/workentry/create", func(c *gin.Context) {
		workEntryCtrl.Post(c)
	})

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
