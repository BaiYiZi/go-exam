package mysql

import (
	"fmt"

	driver_mysql "github.com/BaiYiZi/go-exam/driver/mysql"
)

func Ping(name string) (interface{}, error) {
	connection := driver_mysql.GetDBConnection()

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	response := struct {
		Msg string `json:"msg"`
	}{}

	msg := fmt.Sprintf("Hello %s!, current db is %s", name, "mysql")
	response.Msg = msg

	return response, nil
}
