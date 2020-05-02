package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWithData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Message  string      `json:"message"`
}

type ResponseWithoutData struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

const (
	ERROR   = 10001
	AUTH_ERROR = 10002
	BIND_ERROR = 10003
	SUCCESS = 0
)

func ResultWithData(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, ResponseWithData{
		code,
		data,
		msg,
	})
}

func ResultWithoutData(code int, message string, c *gin.Context) {
	c.JSON(http.StatusOK, ResponseWithoutData{
		code,
		message,
	})
}

func FailErrorBind(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, ResponseWithoutData{
		BIND_ERROR,
		message,
	})
}

func FailErrorAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ResponseWithoutData{
		AUTH_ERROR,
		message,
	})
}

func Ok(c *gin.Context) {
	OkWithMessage("OK", c)
}

func OkWithMessage(message string, c *gin.Context) {
	ResultWithoutData(SUCCESS, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	ResultWithData(SUCCESS, data, "OK", c)
}


func Fail(c *gin.Context) {
	FailWithMessage( "ERROR", c)
}

func FailWithMessage(message string, c *gin.Context) {
	ResultWithoutData(ERROR, message, c)
}
