package handler

import (
	"net/http"

	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-gonic/gin"
)

type MediaHandler struct{}

type GetMediaDetailResp struct {
	BaseResponse
	Data model.Media `json:"data"`
}

func (mh *MediaHandler) GetMediaDetail(c *gin.Context) {
	sid := c.Param("subjectID")

	src, mid, err := service.DecodeID(sid)
	if err != nil {
		rHTTPError(c, errcode.SqidsParseFailed, err.Error(), http.StatusBadRequest)
		return
	}

	ms := new(service.MediaService)
	media, err := ms.MediaDetail(mid, src)
	if err != nil {
		rHTTPError(c, errcode.Failed, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GetMediaDetailResp{}
	resp.Code = 0
	resp.Msg = ""
	resp.Data = *media

	c.JSON(http.StatusOK, resp)
}
