package constants

import (
	"github.com/gin-gonic/gin"
	. "github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/services"
)

func Get(c *gin.Context) {
	data, err := services.GetEnumData()
	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	OkWithData(data, c)
}