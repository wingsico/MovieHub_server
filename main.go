package main

import (
	_ "github.com/wingsico/movie_server/conf"
	_ "github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/routes"
	"os"
)
func main() {
	router := routes.InitRouter()
	router.Run(":" + os.Getenv("SERVER_PORT"))
}


