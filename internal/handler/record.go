package handler

import (
	"net/http"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-gonic/gin"
)

type RecordHandler struct{}

type AddSubjectRecordReq struct {
	Action global.Status `json:"action"`
}

// /api/record/:subjectID POST
func (rh *RecordHandler) AddSubjectRecord(c *gin.Context) {
	sid := c.Param("subjectID")

	_, mid, err := service.DecodeID(sid)
	if err != nil {
		rHTTPError(c, http.StatusBadRequest, err.Error(), errcode.SqidsParseFailed)
		return
	}
	uid := c.MustGet("uid").(uint)
	req := AddSubjectRecordReq{}
	if c.ShouldBindJSON(&req) != nil {
		rHTTPError(c, http.StatusBadRequest, err.Error(), errcode.ParamParseFailed)
		return
	}

	rs := new(service.RecordService)
	if err := rs.RecordSubject(mid, uid, req.Action); err != nil {
		rHTTPError(c, http.StatusInternalServerError, err.Error(), errcode.Failed)
		return
	}

	c.JSON(http.StatusOK, successResp)

}

func (rh *RecordHandler) GetSubjectRecord(c *gin.Context) {
	// rs := new(service.RecordService)
}

func (rh *RecordHandler) RevokeSubjectRecord(c *gin.Context) {
	// rs := new(service.RecordService)
}
