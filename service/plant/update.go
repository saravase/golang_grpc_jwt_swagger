package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdatePlant used to implement client streaming
// To update the plant data into the datastore based on id.
func (server *PlantServer) UpdatePlant(stream pb.PlantService_UpdatePlantServer) error {

	var plants []string
	for {
		plant, serr := stream.Recv()
		if serr == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Updated plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		id, err := server.store.Update(plant)
		if err != nil {
			return err
		}
		plants = append(plants, id)
	}
}
