package router

import (
	"github.com/TravisRoad/goshower/internal/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (*AuthRouter) Register(r *gin.RouterGroup) {
	rt := r.Group("/auth")
	authHandler := &handler.AuthHandler{}

	{
		rt.POST("/login", authHandler.Login)
		rt.POST("/logout", authHandler.Logout)
		rt.GET("/islogin", authHandler.IsLogin)
	}
}
