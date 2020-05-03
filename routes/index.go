package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/handlers/admin"
	"github.com/wingsico/movie_server/handlers/constants"
	"github.com/wingsico/movie_server/handlers/movie/detail"
	"github.com/wingsico/movie_server/handlers/movie/list"
	"github.com/wingsico/movie_server/handlers/upload"
	"github.com/wingsico/movie_server/middlewares"
)

func InitRouter() * gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())
	api := r.Group("/v1/api")
	{
		m := api.Group("/movie")
		{
			l := m.Group("/list")
			{
				l.GET("/search", movie_list.GetBySearch)
				l.GET("/random", movie_list.GetByRandom)
				l.GET("/type", movie_list.GetByType)
				l.POST("/ids", movie_list.GetByIds)
			}

			m.GET("", movie_detail.GetById)

			m.Use(middlewares.AuthMiddleware())
			{
				m.POST("/delete", movie_detail.DeleteByIds)
				m.PUT("/update", movie_detail.Update)
				m.POST("/create", movie_detail.Create)
			}
		}

		a := api.Group("/admin")
		{
			a.POST("/register", admin.Create)
			a.POST("/login", admin.Login)
			a.Use(middlewares.AuthMiddleware()).GET("/info", admin.Get)
		}

		api.GET("/constants", constants.Get)
		u := api.Use(middlewares.AuthMiddleware())
		{
			u.POST("/upload", upload.Upload)
		}

	}

	return r
}