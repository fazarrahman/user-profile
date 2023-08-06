package handler

import (
	"github.com/fazarrahman/user-profile/service"
)

type Server struct {
	Svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{Svc: svc}
}
