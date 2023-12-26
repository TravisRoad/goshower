package router

import (
	"github.com/TravisRoad/goshower/api"
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/admin")
	adminApi := new(api.AdminApi)

	r.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN}))

	r.GET("/user", adminApi.GetUsers)
	r.POST("/user", adminApi.AddUser)
	r.POST("/user/:id", adminApi.UpdateUser)
	r.DELETE("/user/:id", adminApi.DeleteUser)
}
