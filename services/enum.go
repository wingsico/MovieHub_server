package services

import (
	m "github.com/wingsico/movie_server/models"
	"github.com/wingsico/movie_server/response"
)

func GetEnumData() (response.EnumGetResponse,  error) {
	var genres []m.Genre
	var regions []m.Region
	var err error

	if genres, err = (m.Genre{}).GetCommonGenres(); err != nil {
		return response.EnumGetResponse{}, err
	}

	if regions, err = (m.Region{}).GetCommonRegions(); err != nil {
		return response.EnumGetResponse{}, err
	}

	res := response.EnumGetResponse{
		Genres:  genres,
		Regions: regions,
	}

	return res, err
}