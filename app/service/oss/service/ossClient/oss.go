package ossClient

import (
	"libong/common/context"
	"libong/oss/app/service/oss/api"
)

type OssClient interface {
	Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error)
}

type Config struct {
	AliYun *AliYunConf
	Local  *LocalOssConf
}

func NewOss(conf *Config) OssClient {
	if conf.AliYun != nil {
		return newAliYunOss(conf.AliYun)
	} else if conf.Local != nil {
		return newLocalOss(conf.Local)
	} else {
		panic("NewOssClient err")
	}
	return nil
}
