package handler

import (
	"github.com/fazarrahman/user-profile/service"
)

type Server struct {
	Svc service.ServiceInterface
}

func NewServer(svc service.ServiceInterface) *Server {
	return &Server{Svc: svc}
}
