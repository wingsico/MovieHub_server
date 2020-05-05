package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/handlers/admin"
	"github.com/wingsico/movie_server/handlers/constants"
	"github.com/wingsico/movie_server/handlers/movie/detail"
	"github.com/wingsico/movie_server/handlers/movie/list"
	movie_review "github.com/wingsico/movie_server/handlers/movie/review"
	"github.com/wingsico/movie_server/handlers/upload"
	"github.com/wingsico/movie_server/middlewares"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())
	api := r.Group("/v1/api")
	{
		m := api.Group("/movie")
		{
			m.GET("", movie_detail.GetById)

			authorized := m.Group("/")
			authorized.Use(middlewares.AuthMiddleware())
			{
				authorized.POST("/delete", movie_detail.DeleteByIds)
				authorized.PUT("/update", movie_detail.Update)
				authorized.POST("/create", movie_detail.Create)
			}

			l := m.Group("/list")
			{
				l.GET("/search", movie_list.GetBySearch)
				l.GET("/random", movie_list.GetByRandom)
				l.GET("/type", movie_list.GetByType)
				l.POST("/ids", movie_list.GetByIds)
			}

			r := m.Group("/review")
			{
				r.GET("/list", movie_review.GetList)
				r.GET("", movie_review.Get)

				authorized := r.Group("/")
				authorized.Use(middlewares.AuthMiddleware())
				{
					authorized.PUT("/update", movie_review.Update)
					authorized.POST("/delete", movie_review.Delete)
					authorized.POST("/create", movie_review.Create)
				}
			}

		}

		a := api.Group("/admin")
		{
			a.POST("/register", admin.Create)
			a.POST("/login", admin.Login)
			a.GET("/info", middlewares.AuthMiddleware(), admin.Get)
		}

		api.GET("/constants", constants.Get)
		api.POST("/upload", middlewares.AuthMiddleware(), upload.Upload)

	}

	return r
}
