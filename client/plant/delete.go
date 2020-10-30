package plant

import (
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
)

func (client *PlantClient) DeletePlant() {
	dplants := []*pb.PlantID{
		&pb.PlantID{
			Value: "p-2",
		},
		&pb.PlantID{
			Value: "p-3",
		},
	}

	// Request: DeletePlant() - Stream of plant id's

	ds, err := client.cli.DeletePlant(client.ctx)
	if err != nil {
		log.Fatalf("Delete stream creation falied. Reason: %v", err)
	}

	for _, plantId := range dplants {
		if err := ds.Send(plantId); err != nil {
			log.Fatalf("%v.Send(%v) = %v",
				ds, plantId, err)
		}
	}

	deleteRes, err := ds.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			ds, err, nil)
	}
	log.Printf("Deleted Plants Res : %s", deleteRes)

}
