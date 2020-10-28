package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdatePlant(stream pb.PlantService_UpdatePlantServer) error {
	var plants []string

	for {
		plant, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Updated plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		found := true
		for id, _ := range s.plantMap {
			if id == plant.Id {
				s.logger.Printf("Updated plant id : %s", id)
				s.plantMap[id] = plant
				plants = append(plants, id)
				found = false
				break
			}
		}
		if found {
			return status.Errorf(codes.NotFound, "Plant record id %s not found", plant.Id)
		}
	}
}
