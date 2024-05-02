package routers

import (
	"github.com/calvinanto/secret-santa-api-v2/controllers"
	"github.com/gin-gonic/gin"
)

func GameRoutes(router *gin.Engine) {
	router.GET("/games", controllers.GetAllGames)
	router.POST("/new-game")
	router.GET("/game/:game-id", controllers.GetGameStatusById)
	router.GET("/game/:game-id/play")
}
