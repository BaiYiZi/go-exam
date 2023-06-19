package api

import (
	"github.com/BaiYiZi/go-exam/service"
	"github.com/BaiYiZi/go-exam/service/mongo"
	"github.com/BaiYiZi/go-exam/service/mysql"
)

func Ping(name string) (interface{}, error) {
	switch service.GetCurrentDBType() {
	case "mysql":
		return mysql.Ping(name)
	case "mongo":
		return mongo.Ping(name)
	}

	return nil, nil
}
