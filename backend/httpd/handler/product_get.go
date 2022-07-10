package handler

import (
	"github.com/gin-gonic/gin"
	"timecalcy/database"
	"timecalcy/models"
)

func ProductGet() func(c *gin.Context) {
	db := database.Init()

	products := []models.Product{}

	db.Find(&products)

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": products,
		})
	}
}
