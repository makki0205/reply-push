package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/reply-push/service"
)

func CheckToken(c *gin.Context) {
	ok := service.Token.Exists(c.GetHeader("token"))
	if ok {
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"err": "token err",
	})
	c.Abort()
}
