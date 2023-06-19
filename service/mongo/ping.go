package mongo

import (
	"context"
	"fmt"

	"github.com/BaiYiZi/go-exam/driver/mongo"
	"github.com/BaiYiZi/go-exam/model"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Ping(name string) (interface{}, error) {
	// Get client
	client := mongo.GetDBConnection()

	// Operate DB
	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, nil
	}

	response := model.Ping{}.ResponsePing()
	msg := fmt.Sprintf("Hello %s!, current db is %s", name, "mongo")
	response.Msg = msg

	// Return data
	return response, nil
}
