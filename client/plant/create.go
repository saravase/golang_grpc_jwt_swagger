package plant

import (
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
)

func (client *PlantClient) CreatePlant() {
	plants := []*pb.Plant{
		&pb.Plant{
			Id:          "p-1",
			Name:        "Rose",
			Category:    pb.Plant_FLOWER,
			Price:       100.00,
			Description: "Beautiful Flower",
			User: &pb.User{
				Id:    100,
				Name:  "optimus",
				Email: "optimus@gmail.com",
			},
		},
		&pb.Plant{
			Id:          "p-2",
			Name:        "Apple",
			Category:    pb.Plant_FRUIT,
			Price:       500.00,
			Description: "Sweetest Fruit",
			User: &pb.User{
				Id:    101,
				Name:  "primz",
				Email: "primz@gmail.com",
			},
		},
		&pb.Plant{
			Id:          "p-3",
			Name:        "Neem",
			Category:    pb.Plant_TREE,
			Price:       200.00,
			Description: "Powerful Tree",
			User: &pb.User{
				Id:    101,
				Name:  "primz",
				Email: "primz@gmail.com",
			},
		},
	}

	// Request: CreatePlant() - Stream of plant id's
	cs, err := client.cli.CreatePlant(client.ctx)
	if err != nil {
		log.Fatalf("Create stream creation falied. Reason: %v", err)
	}

	for _, plant := range plants {
		if err := cs.Send(plant); err != nil {
			log.Fatalf("%v.Send(%v) = %v",
				cs, plant.Id, err)
		}
	}

	createRes, err := cs.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			cs, err, nil)
	}
	log.Printf("Create Plants Res : %s", createRes)

}
