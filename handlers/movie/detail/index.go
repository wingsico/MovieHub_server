package movie_detail

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/services"
)

func GetById(c *gin.Context) {
	var err error
	var requestParams request.MovieGetRequest
	var movieResponse response.MovieDetailGetResponse
	if err = c.ShouldBindQuery(&requestParams); err != nil {
		helpers.FailErrorBind("缺少必要的电影id", c)
		return
	}

	if movieResponse, err = services.GetMovie(requestParams); err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(movieResponse, c)
}

func DeleteByIds(c *gin.Context) {
	var err error
	var requestParams request.MovieDeleteRequest

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		helpers.FailErrorBind("缺少必要的电影id列表", c)
		return
	}

	if err = services.DeleteMovie(requestParams); err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.Ok(c)
}

func Update(c *gin.Context) {
	var err error
	var requestParams request.MovieUpdateRequest
	var responseData response.MovieDetailGetResponse

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	if responseData, err = services.UpdateMovie(requestParams); err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(responseData, c)
}

func Create(c *gin.Context) {
	var err error
	var requestParams request.MovieCreateRequest
	var responseData response.MovieDetailGetResponse

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	if responseData, err = services.CreateMovie(requestParams); err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(responseData, c)
}