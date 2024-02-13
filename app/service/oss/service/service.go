package service

import (
	"libong/oss/app/service/oss/conf"
	"libong/oss/app/service/oss/service/ossClient"
)

type Service struct {
	ossClient ossClient.OssClient
}

func New(c *conf.Service) *Service {
	return &Service{
		ossClient: ossClient.NewOss(c.OssClientConfig),
	}
}
