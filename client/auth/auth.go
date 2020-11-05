package auth

import (
	"context"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	cli pb.AuthServiceClient
	ctx context.Context
}

func NewAuthClient(conn *grpc.ClientConn, ctx context.Context) *AuthClient {

	return &AuthClient{
		cli: pb.NewAuthServiceClient(conn),
		ctx: ctx,
	}
}
