package router

import (
	"github.com/BaiYiZi/go-exam/controller"
	"github.com/BaiYiZi/go-exam/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Ping())

	ping := v1.Group("/ping")
	ping.POST("/", controller.Ping)
}
