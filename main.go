package main

import (
	_ "github.com/wingsico/movie_server/conf"
	_ "github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/routes"
	"github.com/wingsico/movie_server/services"
	"os"
)
func main() {

	services.InitialModels()

	services.InitialKeys()

	router := routes.InitRouter()
	router.Run(":" + os.Getenv("SERVER_PORT"))
}


