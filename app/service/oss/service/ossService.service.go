package service

import (
	"libong/common/context"
	"libong/oss/app/service/oss/api"
)

func (s *Service) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	return s.ossClient.Upload(ctx, req)
}
