package handler

import (
	"net/http"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary      login
// @Description  login
// @Tags         auth
// @Accept       json
// @Produce      json
// @param        req  body  LoginReq  true  "login info"
// @Success      200  {object}  AuthResponse
// @Failure      500  {object}  BaseResponse
// @Router       /auth/login [post]
func (aa *AuthHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	if uinfo, ok := session.Get(global.USER_INFO_KEY).(map[string]interface{}); ok {
		global.Logger.Info("already login", zap.String("username", uinfo["username"].(string)))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 0,
			"data": uinfo,
			"msg":  "success",
		})
		return
	}

	req := LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": errcode.LoginFailed,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid request", zap.Error(err))
		return
	}

	as := new(service.AuthService)
	user, err := as.Login(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.UsernameOrPwd,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid login auth", zap.String("username", req.Username), zap.Error(err))
		return
	}

	userInfo := model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	session.Set(global.USER_INFO_KEY, userInfo)
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errcode.SessionSave,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": userInfo,
		"msg":  "success",
	})
}

// Logout godoc
// @Summary      logout
// @Description  logout
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  BaseResponse
// @Failure      500  {object}  BaseResponse
// @Router       /auth/logout [post]
func (aa *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1, Path: "/"})
	if err := session.Save(); err != nil {
		rHTTPError(c, errcode.SessionSave, err.Error(), http.StatusInternalServerError)
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, BaseResponse{
		Code: 0,
		Msg:  "success",
	})
}

// IsLogin godoc
// @Summary      IsLogin
// @Description  Islogin
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  AuthResponse
// @Failure      401  {object}  BaseResponse
// @Router       /auth/islogin [get]
func (aa *AuthHandler) IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	uinfo, ok := session.Get(global.USER_INFO_KEY).(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.NotLogin,
			"msg":  "not login",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": uinfo,
	})

}
