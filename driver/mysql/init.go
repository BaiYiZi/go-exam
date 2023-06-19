package mysql

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetDBConnection() *gorm.DB {
	return connection
}

type config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

func init() {
	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := path.Join(p, "config", "driver", "mysql", "config.yaml")
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Charset)
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}

	connection = db
}
