package movie_review

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/services"
	"log"
)

func Get(c *gin.Context) {
	var params request.ReviewGetRequest
	var err error
	var res response.ReviewResponse

	if err := c.ShouldBindQuery(&params); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	res, err = services.GetReview(params)

	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(res, c)
}

func GetList(c *gin.Context) {
	params := request.ReviewGetListRequest {
		Start: 0,
		Limit: 20,
		Sort: "useful_count",
	}
	var err error
	var res response.ReviewListResponse

	if err := c.ShouldBindQuery(&params); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}


	res, err = services.GetReviewList(params)



	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}


	helpers.OkWithData(res, c)
}

func Delete(c *gin.Context) {
	var params request.ReviewDeleteRequest
	var err error

	if err := c.ShouldBind(&params); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	err = services.DeleteReview(params)

	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.Ok(c)
}

func Update(c *gin.Context) {
	var params request.ReviewUpdateRequest
	var err error
	var res response.ReviewResponse

	log.Print(c.Request.Body)

	if err := c.ShouldBindJSON(&params); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	res, err = services.UpdateReview(params)

	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(res, c)
}

func Create(c *gin.Context) {
	var params request.ReviewCreateRequest
	var err error
	var res response.ReviewResponse

	if err := c.ShouldBind(&params); err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	res, err = services.CreateReview(params)

	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(res, c)
}

