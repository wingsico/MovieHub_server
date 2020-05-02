package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/errors"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/request"
	response "github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/services"
)

func Login(c *gin.Context) {
	var requestParams request.AdminLoginRequest
	var res response.AdminLoginResponse
	var err error

	if err := c.ShouldBindJSON(&requestParams); err != nil {
		helpers.FailErrorBind(errors.ErrBind.Error(), c)
		return
	}

	if res, err = services.LoginAdmin(requestParams); err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(res, c)
}