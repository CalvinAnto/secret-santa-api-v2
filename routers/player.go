package routers

import "github.com/gin-gonic/gin"

func PlayerRoutes(router *gin.Engine) {
	router.GET("/player/:player-id")
	router.POST("/player/:player-id")
}
