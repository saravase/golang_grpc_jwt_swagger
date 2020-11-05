package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"github.com/saravase/golang_grpc_jwt_swagger/service/auth"
	"github.com/saravase/golang_grpc_jwt_swagger/service/plant"
	"github.com/saravase/golang_grpc_jwt_swagger/service/user"
	"google.golang.org/grpc"
)

const (
	port       = ":9090"
	tkDuration = 5 * time.Minute
)

func getAccessibleRoles() map[string][]string {

	const plantServicePath = "/pb.AuthService/"

	return map[string][]string{
		"/pb.AuthService/CreatePlant": {"admin"},
		"/pb.AuthService/UpdatePlant": {"admin"},
		"/pb.AuthService/DeletePlant": {"admin"},
		"/pb.AuthService/GetPlants":   {"admin", "user"},
		"/pb.AuthService/GetPlant":    {"admin", "user"},
	}
}

func main() {

	logger := log.New(os.Stdout, "plant-grpc-api", log.LstdFlags)

	// Load .env data
	godotenv.Load()

	logger.Printf("Roles Data ===> ", getAccessibleRoles()["/pb.AuthService/GetPlant"])

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Listening port %s failed", port)
	}

	authProps := auth.NewAuth(os.Getenv("SECRET_KEY"), tkDuration)

	interceptor := auth.NewAuthInterceptor(authProps, getAccessibleRoles())

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	// admin store
	ustore := user.NewInMemoryUserStore()

	// auth server
	as := auth.NewAuthServer(ustore, authProps, logger)
	pb.RegisterAuthServiceServer(s, as)

	// plant server
	pstore := plant.NewInMemoryPlantStore()
	ps := plant.NewPlantServer(logger, pstore)
	pb.RegisterPlantServiceServer(s, ps)

	log.Printf("gRPC server listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serving port %s failed", port)
	}

}
