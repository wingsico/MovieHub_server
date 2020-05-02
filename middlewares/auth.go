package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			helpers.FailErrorAuth(err.Error(), c)
			c.Abort()
			return
		}

		c.Next()
	}
}