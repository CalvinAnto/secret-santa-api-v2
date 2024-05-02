package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		status := c.Writer.Status()

		log.Println(status)

	}
}
