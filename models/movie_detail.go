package models

import (
	. "github.com/wingsico/movie_server/db"
	"time"
)

type MovieDetail struct {
	Id int32 `json:"id" binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
	Title string `json:"title" binding:"required" gorm:"COLUMN:title"`
	Cover string `json:"cover" binding:"required" gorm:"COLUMN:cover"`
	OriginTitle string `json:"origin_title" gorm:"COLUMN:origin_title"`
	DoubanRating float32 `json:"douban_rating" binding:"required" gorm:"COLUMN:douban_rating"`
	PubYear int `json:"pub_year" gorm:"COLUMN:pub_year"`
	Genres []Genre `gorm:"many2many:movie_genre;foreignkey:Id;jointable_foreignkey:movie_id" json:"genres"`
	Regions []Region `gorm:"many2many:movie_region;foreignkey:Id;jointable_foreignkey:movie_id" json:"regions"`
	Directors []Director `gorm:"many2many:movie_director;jointable_foreignkey:movie_id" json:"directors"`
	Writers []Writer `gorm:"many2many:movie_writer;foreignkey:Id;jointable_foreignkey:movie_id" json:"writers"`
	Actors []Actor `gorm:"many2many:movie_actor;foreignkey:Id;jointable_foreignkey:movie_id" json:"actors"`
	DoubanRatingCount int `json:"douban_rating_count" binding:"required" gorm:"COLUMN:douban_rating_count"`
	DoubanId	string `json:"douban_id" binding:"required" gorm:"COLUMN:douban_id"`
	DoubanStarCount float32 `json:"douban_star_count" gorm:"COLUMN:douban_star_count"`
	Lang string `json:"lang" gorm:"COLUMN:lang"`
	IMDbId string `json:"imdb_id" gorm:"COLUMN:imdb_id"`
	ReleaseDate *time.Time `json:"release_date" gorm:"COLUMN:release_date"`
	DoubanSummary string `json:"douban_summary" gorm:"COLUMN:douban_summary"`
	IMDbSummary string `json:"imdb_summary" gorm:"COLUMN:imdb_summary"`
	IMDbRating float32 `json:"imdb_rating" binding:"required" gorm:"COLUMN:imdb_rating"`
	Duration int `json:"duration" gorm:"COLUMN:duration"`
	Alias string `json:"alias" gorm:"COLUMN:alias"`
}

func (MovieDetail) TableName() string {
	return "movies"
}

func (MovieDetail) Get(movieId int) (movie MovieDetail, err error) {
	err = Db.Table("movies").Model(movie).Where("id = ?", movieId).Preload("Regions").Preload("Directors").Preload("Writers").Preload("Actors").Preload("Genres").Find(&movie).Error
	return
}
