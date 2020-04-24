package handlers
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/wingsico/movie_server/models"
//	"github.com/wingsico/movie_server/pkg/errno"
//	"net/http"
//)
//
//func GetEnum(c *gin.Context) {
//	var err error
//	genres, err := models.Genre{}.GetCommonGenres()
//	regions, err := models.Region{}.GetCommonRegions()
//	code, message := errno.DecodeErr(err)
//	if code > 0 {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    code,
//			"message": message,
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    code,
//			"message": message,
//			"data":    gin.H {
//				"genres": genres,
//				"regions": regions,
//			},
//		})
//	}
//}