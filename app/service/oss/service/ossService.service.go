package service

import (
	"context"
	"libong/oss/app/service/oss/api"
)

func (s *Service) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	return s.Upload(ctx, req)
}
