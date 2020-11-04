package plant

import (
	"context"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// GetPlants is used to implement server streaming
// To fetch all the plants data from the datastore
func (server *PlantServer) GetPlants(in *emptypb.Empty, stream pb.PlantService_GetPlantsServer) error {

	plantMap, err := server.store.FindAll()
	if err != nil {
		return err
	}

	for _, plant := range plantMap {
		stream.Send(plant)
	}
	return status.New(codes.OK, "").Err()
}

// GetPlant is used to implement unary stream
// To fetch the plant data based on plant id from the datastore
func (server *PlantServer) GetPlant(ctx context.Context, in *pb.PlantID) (plant *pb.Plant, err error) {

	plantData, err := server.store.Find(in)
	if err != nil {
		return nil, err
	}

	return plantData, status.New(codes.OK, "").Err()
}
