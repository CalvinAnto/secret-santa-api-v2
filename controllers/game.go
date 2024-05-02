package controllers

import (
	"log"
	"net/http"

	"github.com/calvinanto/secret-santa-api-v2/models"
	"github.com/gin-gonic/gin"
)

type Game models.Game

type NewGameRequest struct {
	Size int `json:"size"`
}

func GetAllGamesHandler(c *gin.Context) {

	games, err := models.GetAllGames()

	if err != nil {
		log.Panic(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

func GetGameStatusById(c *gin.Context) {

	param := c.Param("game-id")

	c.IndentedJSON(http.StatusOK, gin.H{"param": param})

}

func PlayGame(c *gin.Context) {

}

func NewGameHandler(c *gin.Context) {

	var newGameRequest NewGameRequest

	c.BindJSON(&newGameRequest)

	log.Println(newGameRequest.Size)

	if (newGameRequest.Size) <= 1 {
		log.Println("Size must be more than 1")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Size must be more than 1"})
		c.Abort()
		return
	}

	id, err := models.NewGame(newGameRequest.Size)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}
