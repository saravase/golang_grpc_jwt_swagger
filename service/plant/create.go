package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlant(stream pb.PlantService_CreatePlantServer) error {

	var plants []string

	for {
		plant, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Created plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		s.logger.Printf("Created plant id : %s", plant.Id)
		s.plantMap[plant.Id] = plant
		plants = append(plants, plant.Id)
	}

}