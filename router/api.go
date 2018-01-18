package router

import (
	"github.com/gin-gonic/gin"
	c "github.com/makki0205/reply-push/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/refresh", c.Refresh)

	auth := api.Group("")
	auth.Use(c.CheckToken)
	auth.POST("/tweet", c.Tweet)

}
