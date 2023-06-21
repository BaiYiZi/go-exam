package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Default().Print("Visit Ping")

		c.Next()
	}
}
