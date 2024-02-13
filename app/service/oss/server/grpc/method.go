package grpc

import (
	"context"
	qContext "libong/common/context"
	"libong/oss/app/service/oss/api"
)

func (s *Server) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	return s.service.Upload(ctx.(qContext.Context), req)
}
