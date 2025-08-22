package service

import (
	commonRedis "libong/common/redis"
	"libong/oss/app/service/oss/conf"
	"libong/oss/app/service/oss/service/ossClient"
)

type Service struct {
	ossClient ossClient.OssClient
}

func New(c *conf.Service) *Service {
	if commonRedis.RedisClient == nil {
		panic("RedisClient not init")
	}
	return &Service{
		ossClient: ossClient.NewOss(c.OssClientConfig),
	}
}
