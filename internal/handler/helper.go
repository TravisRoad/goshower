package handler

import (
	"strconv"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-gonic/gin"
)

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func getPageAndSize(c *gin.Context) (page int, size int) {
	p, ok := c.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(p)
	}
	s, ok := c.GetQuery("size")
	if ok {
		size, _ = strconv.Atoi(s)
	}
	if page <= 0 {
		page = 1
	}
	size = clamp(size, 5, 100)
	return
}

func getSourceType(c *gin.Context) (st service.SourceType) {
	// source := c.Query("source")
	// t := c.Query("type")

	return service.SourceType{
		Source: global.SourceBangumi,
		Type:   global.TypeAnime,
	}
}

func NewHTTPError(c *gin.Context, code int, msg string, statHTTP int) {
	er := BaseResponse{
		Code: code,
		Msg:  msg,
	}
	c.JSON(statHTTP, er)
}
