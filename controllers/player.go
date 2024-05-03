package controllers

import (
	"log"
	"net/http"

	"github.com/calvinanto/secret-santa-api-v2/models"
	"github.com/gin-gonic/gin"
)

type PlayerInfo models.PlayerInfo

type GetPlayerByIdRequest struct {
	PlayerId string `json:"player-id"`
}

func GetPlayerByIdHandler(c *gin.Context) {

	playerId := c.Param("player-id")

	if playerId == "" {
		log.Println("player-id null")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "player-id must be filled"})
		return
	}

	playerInfo, err := models.GetPlayerById(playerId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, playerInfo)

}
