package models

import (
	"github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/types"
	"time"
)

type Review struct {
	Id          int32      `gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
	Title       string     `gorm:"COLUMN:title"`
	Content     string     `gorm:"COLUMN:content"`
	CreatedAt   *time.Time `gorm:"COLUMN:created_at"`
	Source      string     `gorm:"COLUMN:source"`
	UsefulCount int        `gorm:"COLUMN:useful_count"`
	Author      string     `gorm:"COLUMN:author"`
	Rating      float32    `gorm:"COLUMN:rating"`
	SubjectId   int32      `gorm:"COLUMN:subject_id"`
}

func (Review) TableName() string {
	return "reviews"
}

func (Review) Get(id int32) (r Review, err error) {
	err = db.Db.Model(&r).Where("id = ?", id).First(&r).Error
	return
}

func (Review) GetList(rules types.ReviewListGetRules) (rs []Review, total int, err error) {
	var sort = "useful_count"
	ex := db.Db.Model(&rs).Where("subject_id = ?", rules.Id)
	if rules.Sort != "" {
		sort = rules.Sort
	}

	err = ex.Order(sort + " DESC").Offset(rules.Start).Limit(rules.Limit).Find(&rs).Limit(-1).Offset(-1).Count(&total).Error
	return
}

func (r *Review) Create() (err error) {
	err = db.Db.Create(&r).Error
	return
}

func (r *Review) Update() (err error) {
	err = db.Db.Save(&r).Error
	return
}

func (Review) Delete(ids []int32) (err error) {
	err = db.Db.Table("reviews").Where("id in (?)", ids).Delete(Review{}).Error
	return
}
