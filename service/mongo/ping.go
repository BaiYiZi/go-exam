package mongo

import (
	"context"
	"fmt"

	driver_mongo "github.com/BaiYiZi/go-exam/driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Ping(name string) (interface{}, error) {
	client := driver_mongo.GetDBConnection()

	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	response := struct {
		Msg string `json:"msg"`
	}{}

	msg := fmt.Sprintf("Hello %s!, current db is %s", name, "mongo")
	response.Msg = msg

	return response, nil
}
