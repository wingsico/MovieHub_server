package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/services"
)

func Upload(c *gin.Context) {
	file, fileHeader , err := c.Request.FormFile("file")
	if file != nil {
		defer file.Close()
	}

	if err != nil {
		helpers.FailErrorBind(err.Error(), c)
		return
	}

	url, err := services.UploadFile(file, fileHeader.Filename)
	if err != nil {
		helpers.FailWithMessage(err.Error(), c)
		return
	}

	helpers.OkWithData(url, c)
}
