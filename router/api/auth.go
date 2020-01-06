package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/wangweicheng7/Codeartisan/pkg/app"
	"github.com/wangweicheng7/Codeartisan/pkg/e"
	"net/http"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(64)"`
}

func GetAuth(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}
	valid := validation.Validation{}
	username := c.Query("username")
	password := c.Query("password")

	auth := Auth{
		Username: username,
		Password: password,
	}

	ok, _ := valid.Valid(&auth)

	if !ok {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	return
}
