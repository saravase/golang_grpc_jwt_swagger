package plant

import (
	"io"
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// GetPlants used to call GetPlants gRPC service method
func (client *PlantClient) GetPlants() {
	gs, err := client.cli.GetPlants(client.ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Get stream creation failed. Reason : %v", err)
	}

	for {
		// Receive stream of plant
		plant, err := gs.Recv()
		if err == io.EOF {
			log.Printf("Plant records are received...")
			break
		}
		log.Printf("Plant : %v", plant)
	}
}

// GetPlant used to call GetPlant gRPC service method
func (client *PlantClient) GetPlant() {
	plant, err := client.cli.GetPlant(client.ctx, &pb.PlantID{
		Value: "p-1",
	})

	if err != nil {
		log.Fatalf("[ERROR] While find plant %v", err)
	}

	log.Printf("Plant %s: %v", plant.Id, plant)
}
