package helpers

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/wingsico/movie_server/errors"
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
			return si, errors.ErrIdType
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
	if movies == nil || len(movies) == 0 {
		return make([]response.MovieBriefGetResponse, 0)
	}

	for _, m := range movies {
		res = append(res, GetMovieBriefResponse(m))
	}
	return
}

func GetMovieDetailResponse(m models.Movie) (movie response.MovieDetailGetResponse) {
	 if m.Genres == nil || len(m.Genres) == 0 {
	 	m.Genres = make([]models.Genre, 0)
	 }
	if m.Regions == nil || len(m.Regions) == 0 {
		m.Regions = make([]models.Region, 0)
	}
	if m.Directors == nil || len(m.Directors) == 0 {
		m.Directors = make([]models.Director, 0)
	}
	if m.Writers == nil || len(m.Writers) == 0 {
		m.Writers = make([]models.Writer, 0)
	}
	if m.Actors == nil || len(m.Actors) == 0 {
		m.Actors = make([]models.Actor, 0)
	}
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

func TransferReview2Response(r models.Review) (rr response.ReviewResponse) {
	rr = response.ReviewResponse{
		Id:          r.Id,
		Title:       r.Title,
		CreatedAt:   r.CreatedAt,
		Content:     r.Content,
		Author:      r.Author,
		Source:      r.Source,
		SubjectId:   r.SubjectId,
		Rating:      r.Rating,
		UsefulCount: r.UsefulCount,
	}
	return rr
}

func TransferReviewList2Response(rs []models.Review) (rrs []response.ReviewResponse) {
	if rs == nil || len(rs) == 0 {
		return make([]response.ReviewResponse, 0)
	}
	for _, m := range rs {
		rrs = append(rrs, TransferReview2Response(m))
	}
	return
}


