package controller

import (
	"fmt"
	"net/http"

	"github.com/BaiYiZi/go-exam/model"
	service "github.com/BaiYiZi/go-exam/service/api"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// Receive JSON parameters
	req := model.Ping{}.RequestPing()
	if err := c.BindJSON(req); err != nil {
		fmt.Println("error", err)
		return
	}

	// Verify parameters
	if ok, err := req.VerifyParameters(); !ok {
		fmt.Println("error", err)
		return
	}

	// Get response through call service
	response, err := service.Ping(req.Name)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	c.JSON(http.StatusOK, response)
}
