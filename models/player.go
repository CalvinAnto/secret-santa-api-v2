package models

import (
	"log"
	"strings"

	"github.com/calvinanto/secret-santa-api-v2/database"
	"github.com/google/uuid"
)

type Player struct {
	ID         string
	Name       string
	Wishlist   string
	Taken      bool
	ReceiverId string
	GameId     string
}

type PlayerInfo struct {
	Name             string
	Wishlist         string
	ReceiverName     string
	ReceiverWishlist string
}

func NewPlayer(gameId string, size int) error {
	db := database.GetDB()

	query := "INSERT INTO player(id, game_id) VALUES "

	vals := []interface{}{}

	for range size {
		query += "(?, ?),"
		id := uuid.New().String()
		log.Printf("Adding player %s to %s\n", id, gameId)
		vals = append(vals, id, gameId)
	}

	query = strings.TrimSuffix(query, ",")

	stmt, _ := db.Prepare(query)

	_, err := stmt.Exec(vals...)

	if err != nil {
		log.Panic("Error insert player", err)
		return err
	}

	return nil
}

func GetPlayerById(playerId string) (PlayerInfo, error) {

	db := database.GetDB()

	query := `
		SELECT
			player.name as playerName,
			player.wishList as playerWishlist,
			receiver.name as receiverName,
			receiver.wishlist as receiverWishlist
		FROM
			game
			JOIN game_player ON game.id = game_player.game_id
			JOIN player ON player.id = game_player.player_id
			JOIN player as receiver ON receiver.id = player.receiver_id
		WHERE
			player.id = ?
	`

	var playerInfo PlayerInfo

	rows, err := db.Query(query, playerId)

	if err != nil {
		return playerInfo, err
	}

	defer rows.Close()

	rows.Next()

	err = rows.Scan(&playerInfo.Name, &playerInfo.Wishlist, &playerInfo.ReceiverName, &playerInfo.ReceiverWishlist)

	if err != nil {
		return playerInfo, err
	}

	return playerInfo, nil
}
