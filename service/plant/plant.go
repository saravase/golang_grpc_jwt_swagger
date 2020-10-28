package plant

import (
	"github.com/saravase/golang_grpc_jwt_swagger/pb"

	"log"
)

type Server struct {
	logger   *log.Logger
	plantMap map[string]*pb.Plant
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger:   logger,
		plantMap: make(map[string]*pb.Plant),
	}
}
