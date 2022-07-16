package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"timecalcy/httpd/repository"
)

type WorkEntry struct{}

func (*WorkEntry) Post(c *gin.Context) {
	fmt.Println("WorkEntryPost")
	repository.WorkEntryPost(c)
}
