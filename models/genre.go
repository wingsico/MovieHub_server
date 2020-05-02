package models

import (
	. "github.com/wingsico/movie_server/db"
)

type Genre struct {
	Id int32	`json:"id,string"`
	Name string `json:"name"`
}

func (Genre) GetCommonGenres() (genres []Genre, err error) {
	commons := []string {"剧情","喜剧","动作","爱情","科幻","动画","悬疑","惊悚","恐怖","犯罪","同性","音乐","歌舞","传记","历史","战争","西部","奇幻","冒险","灾难","武侠","情色"}
	err = Db.Model(genres).Where("name in (?)", commons).Find(&genres).Error
	return
}

func (Genre) GetByName(name string) (genre Genre, err error) {
	err = Db.Model(genre).Where("name = ?", name).First(&genre).Error
	genre.Name = name
	return
}

func (Genre) GetListByNames(names []string) (genres []Genre) {
	var genre Genre
	genres = []Genre{}
	for _, name := range names {
		genre, _ = Genre{}.GetByName(name)
		genres = append(genres, genre)
	}
	return
}
