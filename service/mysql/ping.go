package mysql

import (
	"fmt"

	driver_mysql "github.com/BaiYiZi/go-exam/driver/mysql"
	"github.com/BaiYiZi/go-exam/model"
)

func Ping(name string) (interface{}, error) {
	// Get connection
	connection := driver_mysql.GetDBConnection()

	// Operate DB
	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	response := model.Ping{}.ResponsePing()
	msg := fmt.Sprintf("Hello %s!, current db is %s", name, "mysql")
	response.Msg = msg

	// Return data
	return response, nil
}
