package handler

import (
	"net/http"
	"strconv"

	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct{}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type GetUsersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Page  int              `json:"page"`
		Size  int              `json:"size"`
		Total int64            `json:"total"`
		Users []model.UserInfo `json:"users"`
	} `json:"data"`
}

// GetUsers godoc
// @Summary      get users list
// @Description  only admin can use this GetUsers
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  GetUsersResponse
// @Failure      500  {object}  BaseResponse
// @Router       /user [get]
func (ah *AdminHandler) GetUsers(c *gin.Context) {
	us := new(service.UserService)

	page, size := getPageAndSize(c)

	users, total, err := us.GetUsers(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.GetUsersFailed,
			"msg":  err.Error(),
		})
		return
	}

	res := GetUsersResponse{}
	res.Code = http.StatusOK
	res.Data.Page = page
	res.Data.Size = size
	res.Data.Total = total
	var usersInfos []model.UserInfo
	for _, u := range users {
		usersInfos = append(usersInfos, model.UserInfo{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	res.Data.Users = usersInfos

	c.JSON(http.StatusOK, res)
}

// UpdateUser godoc
// @Summary      update user info
// @Description  only admin can use this api
// @Tags         admin
// @Accept       json
// @Produce      json
// @param        req  body  UpdateUserRequest  true  "update"
// @Success      200  {object}  BaseResponse
// @Failure      500  {object}  BaseResponse
// @Router       /user/{id} [post]
func (ah *AdminHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.UpdateUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	u := model.User{}
	u.ID = uint(id)
	u.Password = req.Password
	u.Role = req.Role

	us := new(service.UserService)
	if err := us.UpdateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.UpdateUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (ah *AdminHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	us := new(service.UserService)
	if err := us.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.DeleteUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (ah *AdminHandler) AddUser(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	us := new(service.UserService)
	if err := us.AddUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.AddUserFailed,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
