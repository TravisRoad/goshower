package handler

import (
	"strconv"

	"github.com/TravisRoad/goshower/global"
	"github.com/gin-gonic/gin"
)

var (
	successResp = gin.H{
		"code": 0,
		"msg":  "success",
	}
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

func getSourceType(c *gin.Context) (src global.Source, t global.Type) {
	sstr := c.Query("source")
	tstr := c.Query("type")
	var err error
	var x int

	x, err = strconv.Atoi(sstr)
	if err != nil {
		src = global.SourceNil
	} else {
		src = global.Source(x)
	}

	x, err = strconv.Atoi(tstr)
	if err != nil {
		t = global.TypeNil
	} else {
		t = global.Type(x)
	}

	return src, t
}

func rHTTPError(c *gin.Context, code int, msg string, statHTTP int) {
	er := BaseResponse{
		Code: code,
		Msg:  msg,
	}
	c.JSON(statHTTP, er)
}
