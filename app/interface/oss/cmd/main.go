package main

import (
	"libong/common/log"
	commonRedis "libong/common/redis"
	"libong/common/server"
	"libong/oss/app/interface/oss/conf"
	"libong/oss/app/interface/oss/server/http"
	"libong/oss/app/interface/oss/service"
)

func main() {
	config := conf.New()
	log.Init()
	commonRedis.NewClient(config.Service.Dao.Redis)
	svc := service.New(config.Service)
	httpServer := http.NewServer(svc, config.Server.HTTP)
	server.StartWaitingForQuit(httpServer)
}
