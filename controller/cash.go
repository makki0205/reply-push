package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/reply-push/cash"
)

func Refresh(c *gin.Context) {
	cash.ReloadAccount()
	cash.ReloadTokent()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
