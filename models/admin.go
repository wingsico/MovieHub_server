package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/wingsico/movie_server/auth"
	"github.com/wingsico/movie_server/db"
	e "github.com/wingsico/movie_server/errors"
	"regexp"
)

type Admin struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;not null;unique" binding:"required"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required"`
}

func (a *Admin) Encrypt() (err error) {
	a.Password, err = auth.Encrypt(a.Password)
	return
}

func (a *Admin) Validate() error {
	// 验证类型
	if err := validator.New().Struct(a) ; err != nil {
		return e.ErrUserRegisterWrong
	}

	if !CheckUsername(a.Name) {
		return e.ErrUserNameFormatError
	}

	if !CheckPassword(a.Password) {
		return e.ErrUserPasswordFormatError
	}

	if a.Exist() {
		return e.ErrUserNameDuplicate
	}

	return nil
}

func (a *Admin) Create(database *gorm.DB) error {
	if database == nil {
		database = db.Db
	}

	return database.Create(&a).Error
}

func (a *Admin) Compare(pwd string) (err error) {
	err = auth.Compare(a.Password, pwd)
	return
}

func (a *Admin) Exist() bool {
	return !db.Db.Where("name = ?", a.Name).First(&a).RecordNotFound()
}


func GetAdmin(name string) (a Admin, err error) {
	err = db.Db.Where("name = ?", name).First(&a).Error
	return
}

func CheckUsername(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z]([a-zA-Z0-9]){3,15}$", password); !ok {
		return false
	}
	return true
}

func CheckPassword(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z]([a-zA-Z0-9.\\-_]){5,15}$", username); !ok {
		return false
	}
	return true
}


