package setup

import (
	"log/slog"
	"time"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/router"
	"github.com/boj/redistore"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "github.com/TravisRoad/goshower/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const (
	HEALTH_PATH = "/api/health"
	SECRET_KEY  = "iV6pNvjdHVUVc5Q*Wi4S&" // random
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// session
	store, err := redis.NewStore(
		10,
		"tcp",
		global.Config.Redis.Addr,
		global.Config.Redis.Password,
		[]byte(SECRET_KEY),
	)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	_, rs := redis.GetRedisStore(store)
	rs.SetSerializer(redistore.JSONSerializer{})

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: false,
	})
	r.Use(sessions.Sessions("session", store))

	r.Use(
		ginzap.GinzapWithConfig(global.Logger.WithOptions(zap.WithCaller(false)), &ginzap.Config{
			TimeFormat: time.RFC3339,
			UTC:        true,
			SkipPaths:  []string{HEALTH_PATH},
		}),
		ginzap.RecoveryWithZap(global.Logger, true),
	)

	// r.Use(
	// 	gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{HEALTH_PATH}}),
	// 	gin.Recovery(),
	// )

	r.GET(HEALTH_PATH, func(ctx *gin.Context) {})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Register(r)

	return r
}
