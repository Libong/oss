package service

import (
	"libong/common/context"
	commonRedis "libong/common/redis"
	"libong/oss/app/interface/oss/conf"
	ossServiceConf "libong/oss/app/service/oss/conf"
	ossServiceGrpc "libong/oss/app/service/oss/server/grpc"
	ossServiceService "libong/oss/app/service/oss/service"
)

type Service struct {
	ossService *ossServiceGrpc.Server
}

func New(c *conf.Service) (s *Service) {
	if commonRedis.RedisClient == nil {
		panic("RedisClient not init")
	}
	svc := Service{
		ossService: ossServiceGrpc.New(ossServiceService.New(
			&ossServiceConf.Service{
				OssClientConfig: c.OssClientConfig,
			},
		)),
	}
	return &svc
}

// Ping check server ok.
func (s *Service) Ping(ctx context.Context) (err error) {
	return
}
