package request

import (
	"errors"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/models"
	"mime/multipart"
	"strconv"
	"time"
)

type MovieListGetByTypeRequest struct {
	Start    int    `form:"start"`
	Limit    int    `form:"limit"`
	Genres   string `form:"genres"`
	Regions  string `form:"regions"`
	PubYears string `form:"pub_years"`
	Sort     string `form:"sort"`
}

type MovieListGetByRandomRequest struct {
	Start int    `form:"start"`
	Limit int    `form:"limit"`
	Seed  string `form:"seed"`
}

type MovieListGetByIdListRequest struct {
	Ids []string `form:"ids"`
}

type MovieListGetBySearchRequest struct {
	Start    int    `form:"start"`
	Limit    int    `form:"limit"`
	Title    string `form:"q"`
	Genre    string `form:"genre"`
	Region   string `form:"region"`
	PubYears string `form:"pub_years"`
}

type MovieDeleteRequest struct {
	Ids []string `form:"ids" binding:"required"`
}

type MovieGetRequest struct {
	Id string `form:"id" binding:"required"`
}

type MovieCreateRequest struct {
	Title             string   `json:"title" binding:"required"`
	Cover             string   `json:"cover" binding:"required"`
	OriginTitle       string   `json:"origin_title"`
	DoubanRating      float32  `json:"douban_rating"`
	PubYear           int      `json:"pub_year"`
	DoubanRatingCount int      `json:"douban_rating_count"`
	DoubanId          string   `json:"douban_id"`
	Lang              string   `json:"lang"`
	IMDbId            string   `json:"imdb_id"`
	ReleaseDate       *time.Time      `json:"release_date"`
	DoubanSummary     string   `json:"douban_summary" binding:"required"`
	IMDbSummary       string   `json:"imdb_summary"`
	IMDbRating        float32  `json:"imdb_rating"`
	IMDbRatingCount   int      `json:"imdb_rating_count"`
	Duration          int      `json:"duration" binding:"required"`
	Alias             string   `json:"alias"`
	Directors         []string `json:"directors" binding:"required"`
	Actors            []string `json:"actors" binding:"required"`
	Writers           []string `json:"writers"`
	Genres            []string `json:"genres" binding:"required"`
	Regions           []string `json:"regions" binding:"required"`
}

type MovieUpdateRequest struct {
	MovieCreateRequest
	Id string `json:"id" binding:"required"`
}

func (r *MovieCreateRequest) Transfer2Movie() models.Movie {
	model := models.Movie{
		Title:             r.Title,
		Cover:             r.Cover,
		OriginTitle:       r.OriginTitle,
		DoubanRatingCount: r.DoubanRatingCount,
		DoubanRating:      r.DoubanRating,
		DoubanId:          helpers.NewNullString(r.DoubanId),
		DoubanSummary:     r.DoubanSummary,
		IMDbId:            helpers.NewNullString(r.DoubanId),
		IMDbRating:        r.IMDbRating,
		IMDbSummary:       r.IMDbSummary,
		Duration:          r.Duration,
		Alias:             r.Alias,
		ReleaseDate:       r.ReleaseDate,
		PubYear:           r.PubYear,
		Lang:              r.Lang,
		IMDbRatingCount:   r.IMDbRatingCount,
		Genres:            models.Genre{}.GetListByNames(r.Genres),
		Regions:           models.Region{}.GetListByNames(r.Regions),
		Directors:         models.Director{}.GetListByNames(r.Directors),
		Actors:            models.Actor{}.GetListByNames(r.Actors),
		Writers:           models.Writer{}.GetListByNames(r.Writers),
	}
	return model
}

func (r *MovieUpdateRequest) Transfer2Movie() models.Movie {
	id, _ := strconv.Atoi(r.Id)

	model := models.Movie{
		Id:                int32(id),
		Title:             r.Title,
		Cover:             r.Cover,
		OriginTitle:       r.OriginTitle,
		DoubanRatingCount: r.DoubanRatingCount,
		DoubanRating:      r.DoubanRating,
		DoubanId:          helpers.NewNullString(r.DoubanId),
		DoubanSummary:     r.DoubanSummary,
		IMDbId:            helpers.NewNullString(r.DoubanId),
		IMDbRating:        r.IMDbRating,
		IMDbSummary:       r.IMDbSummary,
		IMDbRatingCount:   r.IMDbRatingCount,
		Duration:          r.Duration,
		Alias:             r.Alias,
		ReleaseDate:       r.ReleaseDate,
		PubYear:           r.PubYear,
		Lang:              r.Lang,
		Genres:            models.Genre{}.GetListByNames(r.Genres),
		Regions:           models.Region{}.GetListByNames(r.Regions),
		Directors:         models.Director{}.GetListByNames(r.Directors),
		Actors:            models.Actor{}.GetListByNames(r.Actors),
		Writers:           models.Writer{}.GetListByNames(r.Writers),
	}
	return model
}

type AdminCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Key      string `json:"key" binding:"required"`
}

func (r *AdminCreateRequest) CheckParam() error {
	if r.Name == "" {
		return errors.New("用户名不能为空")
	}

	if r.Password == "" {
		return errors.New("密码不能为空")
	}

	if r.Key == "" {
		return errors.New("激活码不能为空")
	}

	return nil
}

type AdminLoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UploadRequest struct {
	Name string         `form:"name" binding:"required"`
	File multipart.File `form:"file" binding:"required"`
}
