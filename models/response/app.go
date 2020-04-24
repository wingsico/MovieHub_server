package response

import . "github.com/wingsico/movie_server/models"

type EnumData struct {
	Genres []Genre `json:"genres"`
	Regions []Region `json:"regions"`
}

type MovieListData struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Start int `json:"start"`
	Subjects []Movie `json:"subjects"`
}


