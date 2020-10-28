package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeletePlant(stream pb.PlantService_DeletePlantServer) error {
	var plants []string

	for {
		plantId, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Deleted plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		found := true
		for id, _ := range s.plantMap {
			if id == plantId.Value {
				delete(s.plantMap, id)
				plants = append(plants, id)
				found = false
				break
			}
		}

		if found {
			return status.Errorf(codes.NotFound, "Plant record id %s not found", plantId.Value)
		}
	}
}
