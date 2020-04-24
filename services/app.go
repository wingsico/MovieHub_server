package services

import (
	"github.com/wingsico/movie_server/enum"
	"github.com/wingsico/movie_server/helpers"
	m "github.com/wingsico/movie_server/models"
	req "github.com/wingsico/movie_server/models/request"
	res "github.com/wingsico/movie_server/models/response"
	"strconv"
)

func GetEnumData() (data res.EnumData, err error) {
	genres, err := m.Genre{}.GetCommonGenres()
	regions, err := m.Region{}.GetCommonRegions()

	data = res.EnumData{Genres: genres, Regions: regions}
	return
}

func GetListByType(info req.TypeInfo) (movies []m.Movie, total int, err error) {
	genres := helpers.FormatArrayParams(info.Genres, ",")
	regions := helpers.FormatArrayParams(info.Regions, ",")
	pubYears := helpers.FormatArrayParams(info.PubYears, ",")
	orderBy, ok := enum.OrderMap[info.Sort]
	if !ok {
		orderBy = enum.OrderMap[enum.Latest]
	}
	movies, total, err = m.Movie{}.GetList(m.TypeRules{
		Limit:    info.Limit,
		Start:    info.Start,
		Genres:   genres,
		Regions:  regions,
		PubYears: pubYears,
		OrderBy:  orderBy,
	})

	return
}

func GetRandomList(info req.RandomInfo) (movies []m.Movie, total int, err error) {
	orderBy := "RAND(" + info.Seed + ")"

	movies, total, err = m.Movie{}.GetList(m.TypeRules{
		Limit:   info.Limit,
		Start:   info.Start,
		OrderBy: orderBy,
	})
	return
}

func GetListById(info req.IdsInfo) (movies []m.Movie, err error) {
	movies, err = m.Movie{}.GetMoviesByIds(info.Ids)
	return
}

func GetListBySearch(info req.SearchInfo) (movies []m.Movie, total int, err error) {
	movies, total, err = m.Movie{}.GetSearchResult(m.SearchRules{
		Start: info.Start,
		Limit: info.Limit,
		Title: info.Title,
	})
	return
}

func GetDetailById(id string) (movie m.MovieDetail, err error) {
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	movie, err = m.MovieDetail{}.Get(idNumber)
	return
}
