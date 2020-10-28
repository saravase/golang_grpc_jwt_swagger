package plant

import (
	"context"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetPlants(in *emptypb.Empty, stream pb.PlantService_GetPlantsServer) error {

	if len(s.plantMap) == 0 {
		s.logger.Printf("No plant records")
		return status.Error(codes.NotFound, "No plant records")
	}

	for id, plant := range s.plantMap {
		s.logger.Printf(" Sent plant record id : %s", id)
		stream.Send(plant)
	}
	return status.New(codes.OK, "").Err()
}

func (s *Server) GetPlant(ctx context.Context, in *pb.PlantID) (plant *pb.Plant, err error) {

	for id, plant := range s.plantMap {
		s.logger.Printf(" Sent plant record id : %s", id)
		if id == in.Value {
			s.logger.Printf(" Matching plant record id : %s ", id)
			return plant, status.New(codes.OK, "").Err()
		}
	}

	return nil, status.Errorf(codes.NotFound, "Plant record id %s not found", in.Value)
}
