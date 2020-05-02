package models

import (
	"github.com/wingsico/movie_server/db"
)

type Writer struct {
	Id int32 `json:"id,string"`
	Name string `json:"name"`
}


func (Writer) GetByName(name string) (writer Writer, err error) {
	err = db.Db.Model(writer).Where("name = ?", name).First(&writer).Error
	writer.Name = name
	return
}

func (Writer) GetListByNames(names []string) (writers []Writer) {
	var writer Writer
	writers = []Writer{}
	for _, name := range names {
		writer, _ = Writer{}.GetByName(name)
		writers = append(writers, writer)
	}
	return
}