package auth

import (
	"context"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"github.com/saravase/golang_grpc_jwt_swagger/service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	user, err := user.NewUser(req.GetUsername(), req.GetPassword(), req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User creation falied. Reason : %v", err)
	}

	err = server.userStore.Save(user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User registeration falied. Reason : %v", err)
	}

	res := &pb.RegisterResponse{
		Username: user.Username,
	}
	return res, nil
}
