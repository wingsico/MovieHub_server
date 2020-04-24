package models

import (
	"github.com/jinzhu/gorm"
	. "github.com/wingsico/movie_server/db"
)

type Movie struct {
	Id           int32   `json:"id" binding:"required" gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
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

type TypeRules struct {
	Start    int
	Limit    int
	Genres   []int
	PubYears []int
	Regions  []int
	OrderBy  string
}

type SearchRules struct {
	Start int
	Limit int
	Title string
}

func (Movie) GetMoviesByIds(ids []int) (movies []Movie, err error) {
	err = Db.Model(movies).Where("id in (?)", ids).Find(&movies).Error
	return
}

func (Movie) GetList(rules TypeRules) (movies []Movie, total int, err error) {
	ex := Db.Table("movies").Model(movies)
	if len(rules.Genres) != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_genre").Select("movie_id").Where("id in (?)", rules.Genres).QueryExpr())
	}
	if len(rules.Regions) != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_region").Select("movie_id").Where("id in (?)", rules.Regions).QueryExpr())
	}
	if len(rules.PubYears) != 0 {
		if len(rules.PubYears) == 1 {
			ex = ex.Where("pub_year = ?", rules.PubYears[0])
		} else {
			ex = ex.Where("pub_year >= ? and pub_year <= ?", rules.PubYears[0], rules.PubYears[1])
		}
	}
	if rules.OrderBy != "" {
		ex = ex.Order(rules.OrderBy)
	}

	err = ex.Offset(rules.Start).Limit(rules.Limit).Find(&movies).Limit(-1).Offset(-1).Count(&total).Error
	return
}

func (Movie) GetSearchResult(rules SearchRules) (movies []Movie, total int, err error) {
	ex := Db.Model(movies).Where("virtual_keywords like ? ", "%"+rules.Title+"%").Order(
		gorm.Expr("(CASE WHEN virtual_keywords = ? THEN 1 WHEN virtual_keywords like  ? THEN 2 WHEN virtual_keywords like ? THEN 3 WHEN virtual_keywords like ? THEN 4 ELSE 5 END)", rules.Title, rules.Title+"%", "%"+rules.Title+"%", "%"+rules.Title))
	err = ex.Offset(rules.Start).Limit(rules.Limit).Find(&movies).Limit(-1).Offset(-1).Count(&total).Error
	return
}
