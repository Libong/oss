package main

import (
	"libong/common/log"
	commonRedis "libong/common/redis"
	"libong/common/server"
	"libong/common/snowflake"
	"libong/oss/app/service/oss/conf"
	"libong/oss/app/service/oss/server/grpc"
	"libong/oss/app/service/oss/service"
)

// main func1.
func main() {
	config := conf.New()
	log.Init()
	//初始化雪花算法 用于生成id
	snowflake.NewWorker(snowflake.WorkerIDBits, snowflake.DataCenterIDBits)
	commonRedis.NewClient(config.Service.Dao.Redis)
	svc := service.New(config.Service)
	grpcServer := grpc.NewServer(svc, config.Server.GRPC)
	server.StartWaitingForQuit(grpcServer)
}
