package helpers

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/wingsico/movie_server/models"
	"github.com/wingsico/movie_server/response"
	"strconv"
	"strings"
)

func FormatArrayParams(params string, sep string) (arr []int) {
	if params == "" {
		arr = []int{}
	} else {
		strArray := strings.Split(params, sep)
		for _, s := range strArray {
			v, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, v)
			}
		}
	}
	return
}

func SliceAtoi32(sa []string) ([]int32, error) {
	si := make([]int32, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, int32(i))
	}
	return si, nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt32(n int) sql.NullInt32 {
	if n == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: int32(n),
		Valid: true,
	}
}


func TransferMovies2Response(movies []models.Movie) (res []response.MovieBriefGetResponse) {
	for _, m := range movies {
		res = append(res, GetMovieBriefResponse(m))
	}
	return
}

func GetMovieDetailResponse(m models.Movie) (movie response.MovieDetailGetResponse) {
	movie = response.MovieDetailGetResponse{
		Genres:            m.Genres,
		Regions:           m.Regions,
		Directors:         m.Directors,
		Writers:           m.Writers,
		Actors:            m.Actors,
		DoubanRatingCount: m.DoubanRatingCount,
		ReleaseDate:       m.ReleaseDate,
		DoubanSummary:     m.DoubanSummary,
		IMDbSummary:       m.IMDbSummary,
		Alias:             m.Alias,
		Id:                m.Id,
		Title:             m.Title,
		Cover:             m.Cover,
		OriginTitle:       m.OriginTitle,
		DoubanRating:      m.DoubanRating,
		PubYear:           m.PubYear,
		DoubanId:          m.DoubanId.String,
		Lang:              m.Lang,
		IMDbId:            m.IMDbId.String,
		IMDbRating:        m.IMDbRating,
		IMDbRatingCount: m.IMDbRatingCount,
		Duration:          m.Duration,
	}
	return
}

func GetMovieBriefResponse(m models.Movie) (movie response.MovieBriefGetResponse) {
	movie = response.MovieBriefGetResponse{
		Id:           m.Id,
		Title:        m.Title,
		Cover:        m.Cover,
		OriginTitle:  m.OriginTitle,
		DoubanRating: m.DoubanRating,
		PubYear:      m.PubYear,
		DoubanId:     m.DoubanId.String,
		Lang:         m.Lang,
		IMDbId:       m.IMDbId.String,
		IMDbRating:   m.IMDbRating,
		Duration:     m.Duration,
	}
	return
}

func GenerateKey() (k models.Key) {
	k.Value = ksuid.New().String()
	return
}

func BatchGenerateKey(count int) []models.Key {
	var keys []models.Key
	for i := 0; i < count; i++ {
		keys = append(keys, GenerateKey())
	}
	return keys
}

func BatchSaveKey(keys []models.Key) error {
	var err error
	for _, key := range keys {
		if err = key.Create(); err != nil {
			return err
		}
	}
	return nil
}



