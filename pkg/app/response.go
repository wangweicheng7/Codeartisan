package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wangweicheng7/Codeartisan/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	ErrCode int         `json:"error_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 这特么到底是个啥
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		ErrCode: errCode,
		Msg:     e.GetMsg(errCode),
		Data:    data,
	})
	return
}
