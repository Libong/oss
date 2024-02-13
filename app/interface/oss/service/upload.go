package service

import (
	"context"
	ossServiceApi "libong/oss/app/service/oss/api"
)

func (s *Service) Upload(ctx context.Context, data []byte, fileName string, fileType uint32) (interface{}, error) {
	res, err := s.ossService.Upload(ctx, &ossServiceApi.UploadReq{
		Name:               fileName,
		IsKeepOriginalName: false,
		Type:               fileType,
		Data:               data,
	})
	if err != nil {
		return "", err
	}
	return res, nil
}
