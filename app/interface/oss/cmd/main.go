package main

import (
	"libong/common/log"
	"libong/common/server"
	"libong/oss/app/interface/oss/conf"
	"libong/oss/app/interface/oss/server/http"
	"libong/oss/app/interface/oss/service"
)

func main() {
	config := conf.New()
	log.Init()
	svc := service.New(config.Service)
	httpServer := http.NewServer(svc, config.Server.HTTP)
	server.StartWaitingForQuit(httpServer)
}
