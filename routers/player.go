package routers

import (
	"github.com/calvinanto/secret-santa-api-v2/controllers"
	"github.com/gin-gonic/gin"
)

func PlayerRoutes(router *gin.Engine) {
	router.GET("/player/:player-id", controllers.GetPlayerByIdHandler)
	router.POST("/player/:player-id")
}
