package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Page struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	List     interface{} `json:"list"`
}

func PageResponse(c *gin.Context, data Page, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  msg,
	})
}

func Success(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "Ok"
	}
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  msg,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}

func Custom(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}
