package db

import (
	_ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wingsico/movie_server/conf"
	"log"
	"os"
)

var Db *gorm.DB
func init() {
	var err error
	connectionConf := conf.ConnectionConf{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port: os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	Db, err = connectionConf.Connect()

	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		log.Fatalf("database error %v", Db.Error)
	}

	log.Printf("database[%s] connect succeed!", connectionConf.Database)



}
