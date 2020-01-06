package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangweicheng7/Codeartisan/router/api"
)

func Init() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.StaticFS("/export", http.Dir(export))

	// apiv1 := r.Group("/api/v1")

	// apiv1.Use(jwt)

	r.GET("/auth", api.GetAuth)
	return r

}
