package models

import (
	"github.com/jinzhu/gorm"
	"github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/errors"
)

type Key struct {
	gorm.Model
	Value string `gorm:"unique;not null"`
}


func (k *Key) Create() error {
	err := db.Db.Create(&k).Error
	return err
}

func (k *Key) Check(database *gorm.DB) error {
	if database == nil {
		database = db.Db
	}
	ex := database.Where("value = ?", k.Value).First(&k)
	notFound := ex.RecordNotFound()
	if notFound {
		return errors.ErrInvalidKey
	}
	err := ex.Error
	return err
}

func (k *Key) Count(database *gorm.DB) (count int, err error) {
	if database == nil {
		database = db.Db
	}
	err = database.Table("keys").Count(&count).Error
	return
}

func (k *Key) Delete(database *gorm.DB) error {
	if database == nil {
		database = db.Db
	}
	err := database.Unscoped().Delete(&k).Error
	return err
}

func (k *Key) Consume(database *gorm.DB) error {
	if database == nil {
		database = db.Db
	}
	if err := k.Check(database); err != nil {
		return err
	}
	if err := k.Delete(database); err != nil {
		return err
	}
	return nil
}


