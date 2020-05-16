package models

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	. "github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/types"
	"time"
)

type Movie struct {
	Id                int32          `json:"id,string" gorm:"AUTO_INCREMENT;PRIMARY_KEY;TYPE:INT(11);COLUMN:id"`
	Title             string         `json:"title" gorm:"COLUMN:title"`
	Cover             string         `json:"cover" gorm:"COLUMN:cover"`
	OriginTitle       string         `json:"origin_title" gorm:"COLUMN:origin_title"`
	DoubanRating      float32        `json:"douban_rating" gorm:"COLUMN:douban_rating"`
	PubYear           int            `json:"pub_year" gorm:"COLUMN:pub_year"`
	Genres            []Genre        `gorm:"many2many:movie_genre;jointable_foreignkey:movie_id" json:"genres"`
	Regions           []Region       `gorm:"many2many:movie_region;jointable_foreignkey:movie_id" json:"regions"`
	Directors         []Director     `gorm:"many2many:movie_director;jointable_foreignkey:movie_id" json:"directors"`
	Writers           []Writer       `gorm:"many2many:movie_writer;jointable_foreignkey:movie_id" json:"writers"`
	Actors            []Actor        `gorm:"many2many:movie_actor;jointable_foreignkey:movie_id" json:"actors"`
	DoubanRatingCount int            `json:"douban_rating_count" gorm:"COLUMN:douban_rating_count"`
	DoubanId          sql.NullString `json:"douban_id" gorm:"COLUMN:douban_id"`
	Lang              string         `json:"lang" gorm:"COLUMN:lang"`
	IMDbId            sql.NullString `json:"imdb_id" gorm:"COLUMN:imdb_id"`
	IMDbRatingCount   int            `json:"imdb_rating_count" gorm:"COLUMN:imdb_rating_count"`
	ReleaseDate       *time.Time            `json:"release_date" gorm:"COLUMN:release_date"`
	DoubanSummary     string         `json:"douban_summary" gorm:"COLUMN:douban_summary"`
	IMDbSummary       string         `json:"imdb_summary" gorm:"COLUMN:imdb_summary"`
	IMDbRating        float32        `json:"imdb_rating" gorm:"COLUMN:imdb_rating"`
	Duration          int            `json:"duration" gorm:"COLUMN:duration"`
	Alias             string         `json:"alias" gorm:"COLUMN:alias"`
}

func (Movie) TableName() string {
	return "movies"
}

func (m *Movie) Validate() (err error) {
	err = validator.New().Struct(m)
	return
}

func (Movie) Get(movieId int32) (movie Movie, err error) {
	err = Db.Table("movies").Model(movie).Where("id = ?", movieId).Preload("Regions").Preload("Directors").Preload("Writers").Preload("Actors").Preload("Genres").Find(&movie).Error
	return
}

func (Movie) Delete(ids []int32) (err error) {
	err = Db.Table("movies").Where("id in (?)", ids).Delete(Movie{}).Error
	return
}

func (m *Movie) Update() (err error) {
	err = Db.Model(&m).Association("Writers").Replace(m.Writers).Error
	err = Db.Model(&m).Association("Directors").Replace(m.Directors).Error
	err = Db.Model(&m).Association("Regions").Replace(m.Regions).Error
	err = Db.Model(&m).Association("Actors").Replace(m.Actors).Error
	err = Db.Model(&m).Association("Genres").Replace(m.Genres).Error

	err = Db.Save(&m).Error
	return
}

func (m *Movie) Create() (err error) {
	err = Db.Model(&m).Association("Writers").Append(m.Writers).Error
	err = Db.Model(&m).Association("Directors").Append(m.Directors).Error
	err = Db.Model(&m).Association("Regions").Append(m.Regions).Error
	err = Db.Model(&m).Association("Actors").Append(m.Actors).Error
	err = Db.Model(&m).Association("Genres").Append(m.Genres).Error

	err = Db.Save(&m).Error
	return
}


func (Movie) GetMoviesByIds(ids []int32) (movies []Movie, err error) {
	err = Db.Model(movies).Where("id in (?)", ids).Find(&movies).Error
	return
}

func (Movie) GetList(rules types.TypeRules) (movies []Movie, total int, err error) {
	ex := Db.Table("movies").Model(movies)
	if len(rules.Genres) != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_genre").Select("movie_id").Where("genre_id in (?)", rules.Genres).QueryExpr())
	}

	if len(rules.Regions) != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_region").Select("movie_id").Where("region_id in (?)", rules.Regions).QueryExpr())
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

func (Movie) GetSearchResult(rules types.SearchRules) (movies []Movie, total int, err error) {
	//ex := Db.Model(movies).Where("virtual_keywords like ? ", "%"+rules.Title+"%").Order(
	//	gorm.Expr("(CASE WHEN virtual_keywords = ? THEN 1 WHEN virtual_keywords like  ? THEN 2 WHEN virtual_keywords like ? THEN 3 WHEN virtual_keywords like ? THEN 4 ELSE 5 END)", rules.Title, rules.Title+"%", "%"+rules.Title+"%", "%"+rules.Title))

	ex := Db.Model(movies).Where("id in (?)", Db.Model(movies).Select("id").Where("virtual_keywords like ?", "%" + rules.Title + "%").QueryExpr())
	//.Order(
	//		gorm.Expr("(CASE WHEN virtual_keywords = ? THEN 1 WHEN virtual_keywords like  ? THEN 2 WHEN virtual_keywords like ? THEN 3 WHEN virtual_keywords like ? THEN 4 ELSE 5 END)", rules.Title, rules.Title+"%", "%"+rules.Title+"%", "%"+rules.Title))
	if rules.Genre != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_genre").Select("movie_id").Where("genre_id = ?", rules.Genre).QueryExpr())
	}
	if rules.Region != 0 {
		ex = ex.Where("id in (?)", Db.Table("movie_region").Select("movie_id").Where("region_id = ?", rules.Region).QueryExpr())
	}
	if len(rules.PubYears) != 0 {
		if len(rules.PubYears) == 1 {
			ex = ex.Where("pub_year = ?", rules.PubYears[0])
		} else {
			ex = ex.Where("pub_year >= ? and pub_year <= ?", rules.PubYears[0], rules.PubYears[1])
		}
	}
	err = ex.Offset(rules.Start).Limit(rules.Limit).Find(&movies).Limit(-1).Offset(-1).Count(&total).Error
	return
}
