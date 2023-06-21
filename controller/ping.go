package controller

import (
	"fmt"
	"net/http"

	"github.com/BaiYiZi/go-exam/api"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	req := &struct {
		Verify bool   `json:"verify"`
		Name   string `json:"name"`
	}{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !req.Verify {
		c.JSON(http.StatusBadRequest, fmt.Errorf("parameters is incorrect").Error())
		return
	}

	pingServiceFunc, err := api.GetPingServiceFunc()
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	response, err := pingServiceFunc(req.Name)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}
