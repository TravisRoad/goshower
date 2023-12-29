package router

import (
	"github.com/TravisRoad/goshower/internal/handler"
	"github.com/TravisRoad/goshower/internal/middleware"
	"github.com/gin-gonic/gin"
)

type SearchRouter struct{}

func (ar *SearchRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/search")
	searchHandler := new(handler.SearchHandler)

	r.Use(middleware.Auth())

	r.GET("", searchHandler.Search)
}
