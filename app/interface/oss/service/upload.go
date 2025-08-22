package service

import (
	"context"
	ossServiceApi "libong/oss/app/service/oss/api"
)

func (s *Service) Upload(ctx context.Context, req *ossServiceApi.UploadReq) (interface{}, error) {
	res, err := s.ossService.Upload(ctx, req)
	if err != nil {
		return "", err
	}
	return res, nil
}
