package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"hello": "world",
	})

}

func main() {

	router := gin.Default()

	router.GET("/", helloWorld)

	router.Run("localhost:8008")

}
