package main

import (
	"fmt"
	"log/slog"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/setup"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  unlicence
// @license.url   http://unlicense.org/
// @host      localhost:8080
// @BasePath  /api
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	setup.Setup()

	global.Logger.Info("Setup Success")

	r := setup.InitRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Port)); err != nil {
		slog.Error("failed to run server on 0.0.0.0", slog.Int("port", global.Config.Port), err)
	}
}
