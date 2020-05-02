package models

import "github.com/wingsico/movie_server/db"

type Actor struct {
	Id int32 `json:"id,string"`
	Name string `json:"name"`
}

func (Actor) GetByName(name string) (actor Actor, err error) {
	err = db.Db.Model(actor).Where("name = ?", name).First(&actor).Error
	actor.Name = name
	return
}

func (Actor) GetListByNames(names []string) (actors []Actor) {
	var actor Actor
	actors = []Actor{}
	for _, name := range names {
		actor, _ = Actor{}.GetByName(name)
		actors = append(actors, actor)
	}
	return
}
