package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/handlers"
)

func InitRouter() * gin.Engine {
	r := gin.Default()

	s := r.Group("/api/movie")
	{
		s.GET("/enum", handlers.HandleEnumGet)
		s.GET("/list/search", handlers.HandleMovieListGetBySearch)
		s.GET("/list/random", handlers.HandleMovieListGetByRandom)
		s.GET("/list/type", handlers.HandleMovieListGetByType)
		s.POST("/list/ids", handlers.HandleMovieListGetByIds)
		s.GET("/detail", handlers.HandleMovieDetailGetById)
	}

	return r
}