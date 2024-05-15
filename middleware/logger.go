package middleware

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, _ := io.ReadAll(c.Request.Body)
		println(string(body))

		c.Request.Body = io.NopCloser(bytes.NewReader(body))

		c.Next()

		status := c.Writer.Status()

		log.Println(status)

	}
}
