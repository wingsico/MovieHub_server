package models

import "github.com/wingsico/movie_server/db"

type Director struct {
	Id int32 `json:"id,string"`
	Name string `json:"name"`
}

func (Director) GetByName(name string) (director Director, err error) {
	err = db.Db.Model(director).Where("name = ?", name).First(&director).Error
	director.Name = name
	return
}

func (Director) GetListByNames(names []string) (directors []Director) {
	var director Director
	directors = []Director{}
	for _, name := range names {
		director, _ = Director{}.GetByName(name)
		directors = append(directors, director)
	}
	return
}