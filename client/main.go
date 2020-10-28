package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	address = "localhost:9090"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("% connection creataion failed", address)
	}
	defer conn.Close()

	cli := pb.NewPlantServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
	cs, err := cli.CreatePlant(ctx)
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

	GetPlants(cli, ctx)

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
	us, err := cli.UpdatePlant(ctx)
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

	GetPlants(cli, ctx)

	dplants := []*pb.PlantID{
		&pb.PlantID{
			Value: "p-2",
		},
		&pb.PlantID{
			Value: "p-3",
		},
	}

	// Request: DeletePlant() - Stream of plant id's

	ds, err := cli.DeletePlant(ctx)
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

	// Request: GetPlants()
	GetPlants(cli, ctx)

	// Request : GetPlant(Id)
	plant, err := cli.GetPlant(ctx, &pb.PlantID{
		Value: "p-1",
	})

	log.Printf("Plant %s: %v", plant.Id, plant)

}

func GetPlants(cli pb.PlantServiceClient, ctx context.Context) {
	gs, err := cli.GetPlants(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Get stream creation failed. Reason : %v", err)
	}

	for {
		plant, err := gs.Recv()
		if err == io.EOF {
			log.Printf("Plant records are received...")
			break
		}
		log.Printf("Plant : %v", plant)
	}
}
