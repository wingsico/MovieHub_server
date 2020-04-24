package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/enum"
	. "github.com/wingsico/movie_server/helpers"
	req "github.com/wingsico/movie_server/models/request"
	res "github.com/wingsico/movie_server/models/response"
	"github.com/wingsico/movie_server/services"
)

func HandleEnumGet(c *gin.Context) {
	data, err := services.GetEnumData()
	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	OkWithData(data, c)
}

func HandleMovieListGetByType(c *gin.Context) {
	typeInfo := req.TypeInfo{
		Limit: 20,
		Start: 0,
		Sort:  enum.Latest,
	}

	_ = c.ShouldBindQuery(&typeInfo)
	movies, total, err := services.GetListByType(typeInfo)
	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	data := res.MovieListData{
		Limit:    typeInfo.Limit,
		Start:    typeInfo.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}

func HandleMovieListGetByRandom(c *gin.Context) {
	randomInfo := req.RandomInfo{
		Start: 0,
		Limit: 20,
		Seed:  "seed",
	}

	_ = c.ShouldBindQuery(&randomInfo)
	movies, total, err := services.GetRandomList(randomInfo)

	if err != nil {
		FailWithMessage(err.Error(), c)
	}

	data := res.MovieListData{
		Limit:    randomInfo.Limit,
		Start:    randomInfo.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}

func HandleMovieListGetByIds(c *gin.Context) {
	idInfo := req.IdsInfo{Ids: []int{}}

	_ = c.ShouldBindJSON(&idInfo)

	movies, err := services.GetListById(idInfo)
	if err != nil {
		FailWithMessage(err.Error(), c)
		return

	}

	OkWithData(movies, c)
}

func HandleMovieListGetBySearch(c *gin.Context) {
	searchInfo := req.SearchInfo{Start: 0, Limit: 20}

	_ = c.ShouldBindQuery(&searchInfo)
	movies, total, err := services.GetListBySearch(searchInfo)

	if err != nil {
		FailWithMessage(err.Error(), c)
		return

	}

	data := res.MovieListData{
		Limit:    searchInfo.Limit,
		Start:    searchInfo.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}

func HandleMovieDetailGetById(c *gin.Context) {
	idString := c.Query("id")

	data, err := services.GetDetailById(idString)

	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	OkWithData(data, c)
}
