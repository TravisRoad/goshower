package handler

import (
	"net/http"

	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct{}

func (sa *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if len(query) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "query is empty",
			"code": errcode.QueryParseFailed,
		})
		return
	}
	page, size := getPageAndSize(c)

	animeSearcher := service.GetAnimeSearchService()
	animeResult, err := animeSearcher.Search(query, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"code": errcode.Failed,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": animeResult,
	})
}
