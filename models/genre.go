package models

import (
	. "github.com/wingsico/movie_server/db"
)

type Genre struct {
	Id int	`json:"id"`
	Name string `json:"name"`
}

func (Genre) GetCommonGenres() (genres []Genre, err error) {
	commons := []string {"剧情","喜剧","动作","爱情","科幻","动画","悬疑","惊悚","恐怖","犯罪","同性","音乐","歌舞","传记","历史","战争","西部","奇幻","冒险","灾难","武侠","情色"}
	err = Db.Model(genres).Where("name in (?)", commons).Find(&genres).Error
	return
}
