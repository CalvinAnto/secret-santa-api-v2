package main

import (
	"log"
	"net/http"

	"github.com/calvinanto/secret-santa-api-v2/database"
	"github.com/calvinanto/secret-santa-api-v2/middleware"
	"github.com/calvinanto/secret-santa-api-v2/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func HelloWorld(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"hello": "world",
	})

}

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	InitEnv()
	database.InitDatabase()

	router := gin.Default()

	router.Use(middleware.Logger())

	router.GET("/", HelloWorld)

	routers.GameRoutes(router)
	routers.PlayerRoutes(router)

	router.Run("localhost:8008")

}
