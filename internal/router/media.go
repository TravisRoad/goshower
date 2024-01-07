package router

import (
	"github.com/TravisRoad/goshower/internal/handler"
	"github.com/TravisRoad/goshower/internal/middleware"
	"github.com/gin-gonic/gin"
)

type MediaRouter struct{}

func (*MediaRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/media")
	mh := new(handler.MediaHandler)

	r.Use(middleware.Auth())

	r.GET("/:subjectID", mh.GetMediaDetail)
}
