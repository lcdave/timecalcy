package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"timecalcy/database"
	models "timecalcy/models/request"
)

func WorkentryPost() func(c *gin.Context) {
	fmt.Println("WorkentryPost")

	db := database.Init()

	// create workentry
	workentry := models.WorkEntry{
		StartTime: "2020-01-01T00:00:00Z",
		EndTime:   "2020-01-01T00:00:00Z",
		BreakTime: "2020-01-01T00:00:00Z",
	}

	// insert workentry
	res := db.Create(&workentry)

	fmt.Println("result: ", res)

	// return status code 200
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WorkentryPost",
		})
	}
}
