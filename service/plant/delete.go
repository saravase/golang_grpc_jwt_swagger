package plant

import (
	"io"
	"strings"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeletePlant used to implement client streaming
// To delete the plant data into the datastore based on stream of id.
func (server *PlantServer) DeletePlant(stream pb.PlantService_DeletePlantServer) error {

	var plants []string
	for {
		plantId, serr := stream.Recv()
		if serr == io.EOF {
			stream.SendAndClose(&pb.PlantID{
				Value: "Deleted plants : " + strings.Join(plants, ", "),
			})
			return status.New(codes.OK, "").Err()
		}

		id, err := server.store.Delete(plantId)
		if err != nil {
			return err
		}
		plants = append(plants, id)
	}
}
