package setup

import "github.com/TravisRoad/goshower/global"

func Setup() {
	InitViper()
	global.DB = initDB()
	global.Logger = initZap()
	global.Redis = initRedis()
	global.Sqids = initSqids()
}
