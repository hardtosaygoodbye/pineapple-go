package controller

import (
	"net/http"
	"pineapple-go/constant"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code constant.ResponseCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

// Success 成功响应
func Success(ctx *gin.Context, data map[string]interface{}) {

	response := response{
		Code: constant.SUCCESS,
		Msg:  "success",
		Data: data,
	}
	setResponse(ctx, http.StatusOK, response)
}

// Error 错误响应
func Error(ctx *gin.Context, code constant.ResponseCode, msg string) {

	response := response{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

// ErrorWithMsg 错误响应
func ErrorWithMsg(ctx *gin.Context, msg string) {
	response := response{
		Code: 400,
		Msg:  msg,
		Data: gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

func setResponse(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
}

//NOTFOUND method not found action
func NOTFOUND(ctx *gin.Context) {
	response := response{
		Code: constant.CODE_404,
		Msg:  constant.GetCodeText(constant.CODE_404),
		Data: gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusNotFound, response)
}

//StatusInternalServerError server 500 error
func StatusInternalServerError(ctx *gin.Context) {
	response := response{
		Code: constant.CODE_500,
		Msg:  constant.GetCodeText(constant.CODE_500),
		Data: gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusInternalServerError, response)
}
