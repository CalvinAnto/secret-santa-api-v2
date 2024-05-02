package models

import (
	"log"
	"strings"

	"github.com/calvinanto/secret-santa-api-v2/database"
	"github.com/google/uuid"
)

type Game struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

func GetAllGames() ([]Game, error) {
	games := []Game{}

	var db = database.GetDB()

	rows, err := db.Query("SELECT * FROM game")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var game Game

		if err := rows.Scan(&game.ID, &game.Size); err != nil {
			return nil, err
		}

		games = append(games, game)

	}

	return games, err
}

func NewGame(size int) (string, error) {

	db := database.GetDB()

	query := "INSERT INTO game VALUES (?, ?)"

	id := uuid.New().String()
	rows, err := db.Query(query, id, size)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	err = NewPlayer(id, size)

	if err != nil {
		return "", err
	}

	log.Println("Created new game " + id)
	return id, nil
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
