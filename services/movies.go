package services

import (
	_const "github.com/wingsico/movie_server/constants"
	"github.com/wingsico/movie_server/helpers"
	m "github.com/wingsico/movie_server/models"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/types"
	"log"
	"strconv"
)

func GetListByType(req request.MovieListGetByTypeRequest) (movieResponses []response.MovieBriefGetResponse, total int, err error) {
	var Movie m.Movie
	genres := helpers.FormatArrayParams(req.Genres, ",")
	regions := helpers.FormatArrayParams(req.Regions, ",")
	pubYears := helpers.FormatArrayParams(req.PubYears, ",")
	orderBy, ok := _const.OrderMap[req.Sort]
	if !ok {
		orderBy = _const.OrderMap[_const.Latest]
	}
	movies, total, err := Movie.GetList(types.TypeRules{
		Limit:    req.Limit,
		Start:    req.Start,
		Genres:   genres,
		Regions:  regions,
		PubYears: pubYears,
		OrderBy:  orderBy,
	})

	movieResponses = helpers.TransferMovies2Response(movies)

	return
}

func GetListByRandom(req request.MovieListGetByRandomRequest) (movieResponses []response.MovieBriefGetResponse, total int, err error) {
	var Movie m.Movie
	orderBy := "RAND('" + req.Seed + "')"
	movies, total, err := Movie.GetList(types.TypeRules{
		Limit:   req.Limit,
		Start:   req.Start,
		OrderBy: orderBy,
	})

	movieResponses = helpers.TransferMovies2Response(movies)
	return
}

func GetListByIdList(req request.MovieListGetByIdListRequest) (movieResponses []response.MovieBriefGetResponse, err error) {
	var Movie m.Movie
	ids, err := helpers.SliceAtoi32(req.Ids)
	movies, err := Movie.GetMoviesByIds(ids)
	movieResponses = helpers.TransferMovies2Response(movies)
	return
}

func GetListBySearch(req request.MovieListGetBySearchRequest) (movieResponses []response.MovieBriefGetResponse, total int, err error) {
	var Movie m.Movie
	if req.Genre == "" {
		req.Genre = "0"
	}
	if req.Region == "" {
		req.Region = "0"
	}
	genre, err := strconv.Atoi(req.Genre)
	region, err := strconv.Atoi(req.Region)
	pubYears := helpers.FormatArrayParams(req.PubYears, ",")
	if err != nil {
		return movieResponses, 0, err
	}
	movies, total, err := Movie.GetSearchResult(types.SearchRules{
		Start: req.Start,
		Limit: req.Limit,
		Title: req.Title,
		Genre: genre,
		Region: region,
		PubYears: pubYears,
	})

	movieResponses = helpers.TransferMovies2Response(movies)
	return
}


func GetMovie(req request.MovieGetRequest) (response.MovieDetailGetResponse, error) {
	var movie m.Movie
	var err error
	var id int
	if id, err = strconv.Atoi(req.Id); err != nil {
		return response.MovieDetailGetResponse{}, err
	}

	if movie, err = movie.Get(int32(id)); err != nil {
		return response.MovieDetailGetResponse{}, err
	}

	return helpers.GetMovieDetailResponse(movie), nil
}

func DeleteMovie(req request.MovieDeleteRequest) error {
	var movie m.Movie
	var err error
	var ids []int32
	if ids, err = helpers.SliceAtoi32(req.Ids); err != nil {
		return err
	}
	err = movie.Delete(ids)
	return err
}

func UpdateMovie(req request.MovieUpdateRequest) (response.MovieDetailGetResponse, error) {
	var movie m.Movie
	var err error
	var res response.MovieDetailGetResponse

	movie = req.Transfer2Movie()
	if err = movie.Validate(); err != nil {
		return res, err
	}

	if err = movie.Update(); err != nil {
		return res, err
	}

	res = helpers.GetMovieDetailResponse(movie)

	return res, err
}

func CreateMovie(req request.MovieCreateRequest) (response.MovieDetailGetResponse, error) {
	var movie m.Movie
	var err error
	var res response.MovieDetailGetResponse

	movie = req.Transfer2Movie()

	log.Print(movie)

	if err = movie.Validate(); err != nil {
		return res, err
	}

	if err = movie.Create(); err != nil {
		return res, err
	}

	res = helpers.GetMovieDetailResponse(movie)

	return res, err
}
