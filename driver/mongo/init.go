package mongo

import (
	"context"
	"fmt"
	"os"
	"path"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

var client *mongo.Client

func GetDBConnection() *mongo.Client {
	return client
}

type config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func init() {
	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := path.Join(p, "config", "driver", "mongo", "config.yaml")
	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	var cfg config
	err = yaml.Unmarshal(dataBytes, &cfg)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	clt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("error", err)
		return
	}

	client = clt
}
