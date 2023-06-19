package main

import (
	"github.com/BaiYiZi/go-exam/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)

	r.Run(":9090")
}
