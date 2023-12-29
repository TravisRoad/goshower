package router

import (
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/handler"
	"github.com/TravisRoad/goshower/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/admin")
	adminHandler := new(handler.AdminHandler)

	r.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN}))

	r.GET("/user", adminHandler.GetUsers)
	r.POST("/user", adminHandler.AddUser)
	r.POST("/user/:id", adminHandler.UpdateUser)
	r.DELETE("/user/:id", adminHandler.DeleteUser)
}
