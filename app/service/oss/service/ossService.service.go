package service

import (
	"libong/common/context"
	"libong/oss/app/service/oss/api"
)

func (s *Service) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	return s.ossClient.Upload(ctx, req)
}
func (s *Service) MakeFileUrl(ctx context.Context, req *api.MakeFileUrlReq) (*api.MakeFileUrlResp, error) {
	fileUrlMap, err := s.ossClient.MakeFileUrl(ctx, req.Keys)
	if err != nil {
		return nil, err
	}
	return &api.MakeFileUrlResp{
		Map: fileUrlMap,
	}, nil
}
