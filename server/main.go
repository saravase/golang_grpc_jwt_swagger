package main

import (
	"log"
	"net"
	"os"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"github.com/saravase/golang_grpc_jwt_swagger/service/plant"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

func main() {

	logger := log.New(os.Stdout, "plant-grpc", log.LstdFlags)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Listening port %s failed", port)
	}

	s := grpc.NewServer()
	ss := plant.NewServer(logger)
	pb.RegisterPlantServiceServer(s, ss)

	log.Printf("gRPC server listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serving port %s failed", port)
	}

}
