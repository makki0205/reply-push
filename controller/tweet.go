package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/reply-push/service"
)

func Tweet(c *gin.Context) {
	tweet := service.Tweet.Set(c.PostForm("body"))
	c.JSON(http.StatusOK, tweet)
}
