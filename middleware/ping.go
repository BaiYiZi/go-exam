package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Visit Ping")

		c.Next()
	}
}
