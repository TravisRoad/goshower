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
		NewHTTPError(
			c,
			http.StatusBadRequest,
			"query is empty",
			errcode.QueryParseFailed,
		)
		return
	}
	page, size := getPageAndSize(c)

	searcher := service.GetSearcher(getSourceType(c))
	result, err := searcher.Search(query, page, size)
	if err != nil {
		NewHTTPError(
			c,
			http.StatusInternalServerError,
			err.Error(),
			errcode.Failed,
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
