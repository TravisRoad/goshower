package router

import (
	"github.com/TravisRoad/goshower/internal/handler"
	"github.com/TravisRoad/goshower/internal/middleware"
	"github.com/gin-gonic/gin"
)

type RecordRouter struct{}

func (*RecordRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/record")
	recordHandler := new(handler.RecordHandler)

	r.Use(middleware.Auth())

	r.POST("/:subjectID", recordHandler.AddSubjectRecord)
	r.GET("/:subjectID", recordHandler.GetSubjectRecord)
	r.DELETE("/:subjectID", recordHandler.RevokeSubjectRecord)
}
