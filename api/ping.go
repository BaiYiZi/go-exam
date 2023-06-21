package api

import (
	"fmt"

	"github.com/BaiYiZi/go-exam/service"
	"github.com/BaiYiZi/go-exam/service/mongo"
	"github.com/BaiYiZi/go-exam/service/mysql"
)

type PingServiceFunc func(name string) (interface{}, error)

func GetPingServiceFunc() (PingServiceFunc, error) {
	switch service.GetCurrentDBType() {
	case "mysql":
		return mysql.Ping, nil

	case "mongo":
		return mongo.Ping, nil

	default:
		return nil, fmt.Errorf("found ping func")
	}
}
