package response

import (
	. "github.com/wingsico/movie_server/models"
	"time"
)

type EnumGetResponse struct {
	Genres  []Genre  `json:"genres"`
	Regions []Region `json:"regions"`
}

type MovieListGetResponse struct {
	Total    int                     `json:"total"`
	Limit    int                     `json:"limit"`
	Start    int                     `json:"start"`
	Subjects []MovieBriefGetResponse `json:"subjects"`
}

type MovieBriefGetResponse struct {
	Id           int32   `json:"id,string" binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
	Title        string  `json:"title" binding:"required" gorm:"COLUMN:title"`
	Cover        string  `json:"cover" binding:"required" gorm:"COLUMN:cover"`
	OriginTitle  string  `json:"origin_title" gorm:"COLUMN:origin_title"`
	DoubanRating float32 `json:"douban_rating" binding:"required" gorm:"COLUMN:douban_rating"`
	PubYear      int     `json:"pub_year" gorm:"COLUMN:pub_year"`
	DoubanId     string  `json:"douban_id" binding:"required" gorm:"COLUMN:douban_id"`
	Lang         string  `json:"lang" gorm:"COLUMN:lang"`
	IMDbId       string  `json:"imdb_id" gorm:"COLUMN:imdb_id"`
	IMDbRating   float32 `json:"imdb_rating" binding:"required" gorm:"COLUMN:imdb_rating"`
	Duration     int     `json:"duration" gorm:"COLUMN:duration"`
}

type MovieDetailGetResponse struct {
	Id                int32      `json:"id,string" binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
	Title             string     `json:"title" binding:"required" gorm:"COLUMN:title"`
	Cover             string     `json:"cover" binding:"required" gorm:"COLUMN:cover"`
	OriginTitle       string     `json:"origin_title" gorm:"COLUMN:origin_title"`
	DoubanRating      float32    `json:"douban_rating" binding:"required" gorm:"COLUMN:douban_rating"`
	PubYear           int        `json:"pub_year" gorm:"COLUMN:pub_year"`
	Genres            []Genre    `gorm:"many2many:movie_genre;jointable_foreignkey:movie_id" json:"genres"`
	Regions           []Region   `gorm:"many2many:movie_region;jointable_foreignkey:movie_id" json:"regions"`
	Directors         []Director `gorm:"many2many:movie_director;jointable_foreignkey:movie_id" json:"directors"`
	Writers           []Writer   `gorm:"many2many:movie_writer;jointable_foreignkey:movie_id" json:"writers"`
	Actors            []Actor    `gorm:"many2many:movie_actor;jointable_foreignkey:movie_id" json:"actors"`
	DoubanRatingCount int        `json:"douban_rating_count" binding:"required" gorm:"COLUMN:douban_rating_count"`
	DoubanId          string     `json:"douban_id" binding:"required" gorm:"COLUMN:douban_id"`
	Lang              string     `json:"lang" gorm:"COLUMN:lang"`
	IMDbId            string     `json:"imdb_id" gorm:"COLUMN:imdb_id"`
	ReleaseDate       *time.Time  `json:"release_date" gorm:"COLUMN:release_date"`
	DoubanSummary     string     `json:"douban_summary" gorm:"COLUMN:douban_summary"`
	IMDbSummary       string     `json:"imdb_summary" gorm:"COLUMN:imdb_summary"`
	IMDbRating        float32    `json:"imdb_rating" binding:"required" gorm:"COLUMN:imdb_rating"`
	Duration          int        `json:"duration" gorm:"COLUMN:duration"`
	Alias             string     `json:"alias" gorm:"COLUMN:alias"`
	IMDbRatingCount   int        `json:"imdb_rating_count" gorm:"COLUMN:imdb_rating_count"`
}

type AdminCreateResponse struct {
	Name string `json:"name"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type AdminInfoResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
