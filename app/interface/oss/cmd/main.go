package main

import (
	"libong/common/log"
	commonRedis "libong/common/redis"
	"libong/common/server"
	"libong/common/snowflake"
	"libong/login/auth"
	"libong/oss/app/interface/oss/conf"
	"libong/oss/app/interface/oss/server/http"
	"libong/oss/app/interface/oss/service"
)

func main() {
	config := conf.New()
	log.Init()
	//初始化雪花算法 用于生成id
	snowflake.NewWorker(snowflake.WorkerIDBits, snowflake.DataCenterIDBits)
	commonRedis.NewClient(config.Service.Dao.Redis)
	svc := service.New(config.Service)
	httpServer := http.NewServer(svc, config.Server.HTTP)
	//初始化授权
	auth.New(&auth.Config{
		RBACService: config.Service.RBACService,
	})
	server.StartWaitingForQuit(httpServer)
}
