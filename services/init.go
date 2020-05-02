package services

import (
	"github.com/wingsico/movie_server/constants"
	"github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/models"
	"log"
)

func InitialKeys() {
	var key models.Key
	var count int
	var err error


	if count, err = key.Count(nil); err != nil {
		panic("激活码初始化失败[Count]: " + err.Error())
		return
	}

	if count < constants.MinKeyCount {
		keys := helpers.BatchGenerateKey(constants.MinKeyCount)
		if err = helpers.BatchSaveKey(keys); err != nil {
			panic("激活码初始化失败[Save]: " + err.Error())
			return
		}
	}

	log.Print("激活码初始化成功")
}

func InitialModels() {
	db.Db.AutoMigrate(&models.Admin{})
	db.Db.AutoMigrate(&models.Key{})
}
