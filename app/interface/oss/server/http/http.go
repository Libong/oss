package http

import (
	"libong/common/server/http"
	"libong/login/auth"
	"libong/oss/app/interface/oss/service"
)

var svc *service.Service

// NewServer .
func NewServer(s *service.Service, c *http.Config) *http.Server {
	server := http.New(c)
	ConfigHttp(s, server)
	return server
}

func ConfigHttp(s *service.Service, server *http.Server) *http.Server {
	svc = s
	authGroup := server.Group("/x/api")
	authGroup.Use(auth.Authorize)
	authGroup.POST("/upload", upload)
	return server
}
