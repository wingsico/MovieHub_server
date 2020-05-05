package services

import (
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/models"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/types"
	"strconv"
)

func GetReview(req request.ReviewGetRequest) (response.ReviewResponse,  error) {
	var r models.Review
	var id int
	var err error
	if id, err = strconv.Atoi(req.Id); err != nil {
		return response.ReviewResponse{}, err
	}

	review, err := r.Get(int32(id))

	return helpers.TransferReview2Response(review), nil
}

func CreateReview(req request.ReviewCreateRequest) (response.ReviewResponse,  error) {
	r, err := req.Transfer2Review()
	if err != nil {
		return response.ReviewResponse{}, err
	}

	if err := r.Create(); err != nil {
		return response.ReviewResponse{}, err
	}

	return helpers.TransferReview2Response(r), nil
}

func UpdateReview(req request.ReviewUpdateRequest) (response.ReviewResponse,  error) {
	r, err := req.Transfer2Review()
	if err != nil {
		return response.ReviewResponse{}, err
	}

	if err := r.Update(); err != nil {
		return response.ReviewResponse{}, err
	}

	return helpers.TransferReview2Response(r), nil
}

func DeleteReview(req request.ReviewDeleteRequest) error {
	var r models.Review
	ids, err := helpers.SliceAtoi32(req.Ids)
	if err != nil {
		return err
	}
	if err := r.Delete(ids); err != nil {
		return err
	}

	return nil
}

func GetReviewList(req request.ReviewGetListRequest) (response.ReviewListResponse,error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return response.ReviewListResponse{}, nil
	}

	subject, err := GetMovie(request.MovieGetRequest{Id: req.Id})


	rule := types.ReviewListGetRules{
		Start: req.Start,
		Limit: req.Limit,
		Id:    id,
		Sort: req.Sort,
	}

	rs, total, err := models.Review{}.GetList(rule)

	if err != nil {
		return response.ReviewListResponse{}, err
	}


	return response.ReviewListResponse{
		Start: req.Start,
		Limit: req.Limit,
		Total: total,
		Reviews: helpers.TransferReviewList2Response(rs),
		Subject: subject,
	}, nil
}