package plant

import (
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
)

func (client *PlantClient) UpdatePlant() {
	uplants := []*pb.Plant{
		&pb.Plant{
			Id:          "p-1",
			Name:        "Rose",
			Category:    pb.Plant_FLOWER,
			Price:       150.00,
			Description: "Beautiful Flower",
			User: &pb.User{
				Id:    101,
				Name:  "primz",
				Email: "primz@gmail.com",
			},
		},
		&pb.Plant{
			Id:          "p-2",
			Name:        "Apple",
			Category:    pb.Plant_FRUIT,
			Price:       700.00,
			Description: "Sweetest Fruit",
			User: &pb.User{
				Id:    100,
				Name:  "optimus",
				Email: "optimus@gmail.com",
			},
		},
	}

	// Request: UpdatePlant() - Stream of plant id's
	us, err := client.cli.UpdatePlant(client.ctx)
	if err != nil {
		log.Fatalf("Update stream creation falied. Reason: %v", err)
	}

	for _, plant := range uplants {
		if err := us.Send(plant); err != nil {
			log.Fatalf("%v.Send(%v) = %v",
				us, plant.Id, err)
		}
	}

	updateRes, err := us.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			us, err, nil)
	}
	log.Printf("Updated Plants Res : %s", updateRes)

}
