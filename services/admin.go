package services

import (
	"github.com/wingsico/movie_server/auth"
	_ "github.com/wingsico/movie_server/conf"
	"github.com/wingsico/movie_server/constants"
	"github.com/wingsico/movie_server/db"
	"github.com/wingsico/movie_server/errors"
	"github.com/wingsico/movie_server/helpers"
	"github.com/wingsico/movie_server/models"
	"github.com/wingsico/movie_server/request"
	"github.com/wingsico/movie_server/response"
	"github.com/wingsico/movie_server/token"
	"log"
)

func CreateAdmin(r request.AdminCreateRequest) error {
	// 检查了用户名密码key存在性和格式
	if err := r.CheckParam(); err != nil {
		return err
	}

	admin := models.Admin{
		Name:     r.Name,
		Password: r.Password,
	}

	// 验证格式
	if err := admin.Validate(); err != nil {
		return err
	}

	// 加密密码
	if err := admin.Encrypt(); err != nil {
		return err
	}

	log.Print(admin)
	// 创建管理员
	// 创建事务，激活码消耗和用户创建必须保持一致
	tx := db.Db.Begin()
	if err := admin.Create(tx); err != nil {
		tx.Rollback()
		return err
	}

	key := models.Key{Value: r.Key}

	if err := key.Consume(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	var count int
	var err error
	if count, err = key.Count(nil); err != nil {
		return err
	}
	if count < constants.MinKeyCount {
		keys := helpers.BatchGenerateKey(constants.MinKeyCount)
		if err = helpers.BatchSaveKey(keys); err != nil {
			log.Print("key生成失败: " + err.Error())
		}
	}

	return nil
}

func LoginAdmin(r request.AdminLoginRequest) (res response.AdminLoginResponse, err error) {
	a, err := models.GetAdmin(r.Name)
	if err != nil {
		return res, errors.ErrUserLoginWrong
	}

	if err := auth.Compare(a.Password, r.Password); err != nil {
		return res, errors.ErrUserLoginWrong
	}

	t, err := token.Sign(token.Context{ID: uint64(a.ID), Username: a.Name }, "")
	if err != nil {
		return res, errors.ErrToken
	}

	return response.AdminLoginResponse{Token: t}, nil
}
