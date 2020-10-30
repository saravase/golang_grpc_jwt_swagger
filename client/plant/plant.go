package plant

import (
	"context"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc"
)

type PlantClient struct {
	cli pb.PlantServiceClient
	ctx context.Context
}

func NewPlantClient(conn *grpc.ClientConn, ctx context.Context) *PlantClient {

	return &PlantClient{
		cli: pb.NewPlantServiceClient(conn),
		ctx: ctx,
	}
}
