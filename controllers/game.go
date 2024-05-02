package controllers

import (
	"log"
	"net/http"

	"github.com/calvinanto/secret-santa-api-v2/database"
	"github.com/calvinanto/secret-santa-api-v2/models"
	"github.com/gin-gonic/gin"
)

type Game models.Game

func GetAllGames(c *gin.Context) {

	games := []Game{}

	var db = database.GetDB()

	rows, err := db.Query("SELECT * FROM game")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var game Game

		if err := rows.Scan(&game.ID, &game.Size); err != nil {
			log.Fatal(err)
			return
		}

		games = append(games, game)

	}

	c.IndentedJSON(http.StatusOK, games)
}

func GetGameStatusById(c *gin.Context) {

	param := c.Param("game-id")

	c.IndentedJSON(http.StatusOK, gin.H{"param": param})

}

func PlayGame(c *gin.Context) {

}
