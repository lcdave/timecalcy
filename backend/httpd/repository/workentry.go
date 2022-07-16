package repository

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timecalcy/database"
	models "timecalcy/models/request"
)

func WorkEntryPost(c *gin.Context) {
	var workEntryRequest models.WorkEntry

	if err := c.ShouldBindJSON(&workEntryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workEntry := models.WorkEntry{
		StartTime: workEntryRequest.StartTime,
		EndTime:   workEntryRequest.EndTime,
		BreakTime: workEntryRequest.BreakTime,
	}

	db := database.Init()

	if err := db.Create(&workEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workEntry})
}
