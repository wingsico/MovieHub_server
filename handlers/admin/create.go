package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/services"
	"log"
)


func Create(c *gin.Context) {
	var r request.AdminCreateRequest
	if err := c.Bind(&r); err != nil {
		helpers.FailErrorBind("请求格式有误或缺少参数", c)
		return
	}
	log.Print(r)
	if err := services.CreateAdmin(r); err!= nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	res := response.AdminCreateResponse{Name: r.Name}

	helpers.OkWithData(res, c)
}

