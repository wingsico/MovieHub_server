package models

import (
	. "github.com/wingsico/movie_server/db"
)

type Region struct {
	Id int32 `json:"id,string"`
	Name string `json:"name"`
}

func (Region) GetCommonRegions() (regions []Region, err error) {
	commons := []string {"中国大陆", "中国台湾", "中国香港", "美国", "英国", "法国", "韩国", "泰国", "印度", "意大利", "俄罗斯", "德国"}
	err = Db.Model(regions).Where("name in (?)", commons).Find(&regions).Error
	return
}

func (Region) GetByName(name string) (region Region, err error) {
	err = Db.Model(region).Where("name = ?", name).First(&region).Error
	region.Name = name
	return
}

func (Region) GetListByNames(names []string) (regions []Region) {
	var region Region
	regions = []Region{}
	for _, name := range names {
		region, _ = Region{}.GetByName(name)
		regions = append(regions, region)
	}
	return
}