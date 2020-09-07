package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    ResposeCode `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, c ResposeCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	respose, _ := json.Marshal(rd)
	ctx.Set("response", string(respose))
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResposeCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	respose, _ := json.Marshal(rd)
	ctx.Set("response", string(respose))
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	respose, _ := json.Marshal(rd)
	ctx.Set("response", string(respose))
	ctx.JSON(http.StatusOK, rd)
}
