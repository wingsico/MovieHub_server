package errors

import "errors"

var (
	ErrUserNameDuplicate = errors.New("注册的用户名已存在")
	ErrBind = errors.New("请求的数据格式有误或缺少必要参数")
	ErrInvalidKey = errors.New("无效的key")
	ErrUserLoginWrong = errors.New("用户名或密码有误")
	ErrUserRegisterWrong = errors.New("用户名或密码格式有误")
	ErrMissingAuthHeader = errors.New("authorization头部不存在")
	ErrToken = errors.New("token生成错误")
	ErrInvalidToken = errors.New("无效的token")
	ErrUserNameFormatError = errors.New("用户名只能是4-16位的英文或数字的组合")
	ErrUserPasswordFormatError = errors.New("密码只能是4-16位的英文或数字的组合")
	ErrIdType = errors.New("id必须是正整数字符串")
	ErrMovieNotExisted = errors.New("该电影不存在")
	ErrReviewNotExisted = errors.New("该影评不存在")
)
