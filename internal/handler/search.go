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
		rHTTPError(
			c,
			http.StatusBadRequest,
			"query is empty",
			errcode.QueryParseFailed,
		)
		return
	}
	page, size := getPageAndSize(c)
	source, t := getSourceType(c)

	searchService := new(service.SearchService)
	result, err := searchService.Search(query, page, size, source, t)
	if err != nil {
		rHTTPError(
			c,
			errcode.Failed,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": result,
	})
}

func (sa *SearchHandler) SearchSourceType(c *gin.Context) {

}
