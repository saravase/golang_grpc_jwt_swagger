package auth

import (
	"context"
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	log.Printf("User Name Details : %s", req.Username)
	user, err := server.userStore.Find(req.Username)
	log.Printf("User Details : %v", user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not exist. Create a new account")
	}

	if user == nil || !user.VerifyPassword(req.Password) {
		return nil, status.Errorf(codes.NotFound, "Incorrect user credentials")
	}

	token, err := server.auth.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Access token creation falied. Reason : %v", err)
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}
