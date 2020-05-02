package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/token"
)

func Get(c *gin.Context) {
	context, _ := token.ParseRequest(c)

	helpers.OkWithData(&response.AdminInfoResponse{
		Id: string(context.ID),
		Name: context.Username,
	}, c)
}