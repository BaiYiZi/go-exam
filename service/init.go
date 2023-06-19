package service

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var db string

func GetCurrentDBType() string {
	return db
}

type config struct {
	DB string `yaml:"db"`
}

func init() {
	p, _ := os.Getwd()
	filePath := path.Join(p, "config", "service", "config.yaml")

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

	db = cfg.DB
}
