package main

import (
	"libong/common/log"
	"libong/common/server"
	"libong/oss/app/service/oss/conf"
	"libong/oss/app/service/oss/server/grpc"
	"libong/oss/app/service/oss/service"
)

// main func1.
func main() {
	config := conf.New()
	log.Init()
	svc := service.New(config.Service)
	grpcServer := grpc.NewServer(svc, config.Server.GRPC)
	server.StartWaitingForQuit(grpcServer)
}
