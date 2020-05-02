package movie_list

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/constants"
	. "github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/services"
)

func GetByType(c *gin.Context) {
	requestParams := request.MovieListGetByTypeRequest{
		Limit: 20,
		Start: 0,
		Sort:  constants.Latest,
	}

	_ = c.ShouldBindQuery(&requestParams)
	movies, total, err := services.GetListByType(requestParams)

	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	data := response.MovieListGetResponse{
		Limit:    requestParams.Limit,
		Start:    requestParams.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}

func GetByRandom(c *gin.Context) {
	requestParams := request.MovieListGetByRandomRequest{
		Start: 0,
		Limit: 20,
		Seed:  "0",
	}

	_ = c.ShouldBindQuery(&requestParams)

	movies, total, err := services.GetListByRandom(requestParams)

	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	data := response.MovieListGetResponse {
		Limit:    requestParams.Limit,
		Start:    requestParams.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}

func GetByIds(c *gin.Context) {
	var err error
	var movies []response.MovieBriefGetResponse
	requestParams := request.MovieListGetByIdListRequest{
		Ids: []string{},
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		FailErrorBind("请求格式有误，缺少 ids", c)
		return
	}

	if movies, err = services.GetListByIdList(requestParams); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	OkWithData(movies, c)
}

func GetBySearch(c *gin.Context) {
	requestParams := request.MovieListGetBySearchRequest{
		Start: 0,
		Limit: 20,
		Title: "",
	}

	if err := c.ShouldBindQuery(&requestParams); err != nil {
		FailErrorBind(err.Error(), c)
		return
	}

	movies, total, err := services.GetListBySearch(requestParams)

	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}

	data := response.MovieListGetResponse {
		Limit:    requestParams.Limit,
		Start:    requestParams.Start,
		Total:    total,
		Subjects: movies,
	}

	OkWithData(data, c)
}
