package middleware

import (
	"fmt"
	"net/http"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/helper"
	"github.com/TravisRoad/goshower/internal/errcode"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func debugUserinfo(s sessions.Session) (map[string]interface{}, error) {
	uinfo := make(map[string]interface{})
	uinfo["id"] = 1
	uinfo["username"] = "admin"
	uinfo["role"] = "admin"

	s.Set(global.USER_INFO_KEY, uinfo)
	if err := s.Save(); err != nil {
		return nil, fmt.Errorf("debugUserinfo failed to save session: %w", err)
	}
	return uinfo, nil
}

func saveUserInfo(uinfo map[string]interface{}, c *gin.Context) {
	c.Set("username", uinfo["username"])
	c.Set("ID", uinfo["id"])
	c.Set("role", uinfo["role"])
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uinfo, ok := session.Get(global.USER_INFO_KEY).(map[string]interface{})
		mode := helper.Mode()

		// if debug mode on
		if debug := c.Query("debug"); !ok && len(debug) != 0 && (mode == global.DEV || mode == global.TEST) {
			userinfo, err := debugUserinfo(session)
			if err != nil {
				global.Logger.Error("failed to debugUserinfo", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.SessionSave,
					"msg":  err.Error(),
				})
				return
			}
			saveUserInfo(userinfo, c)
			c.Next()
			return
		}

		// or just return unauthorized
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": errcode.NotLogin,
				"msg":  "not login",
			})
			return
		}

		saveUserInfo(uinfo, c)
		c.Next()
	}
}
