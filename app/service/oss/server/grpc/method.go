package grpc

import (
	"context"
	qContext "libong/common/context"
	"libong/oss/app/service/oss/api"
)

func (s *Server) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	return s.service.Upload(ctx.(qContext.Context), req)
}
func (s *Server) MakeFileUrl(ctx context.Context, req *api.MakeFileUrlReq) (*api.MakeFileUrlResp, error) {
	return s.service.MakeFileUrl(ctx.(qContext.Context), req)
}
