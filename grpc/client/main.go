package main

import (
	"context"
	"log"
	"time"

	"github.com/saravase/golang_grpc_jwt_swagger/client/plant"
	"google.golang.org/grpc"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pclient := plant.NewPlantClient(conn, ctx)
	pclient.CreatePlant()
	pclient.GetAllPlant()
	pclient.UpdatePlant()
	pclient.GetAllPlant()
	pclient.DeletePlant()
	pclient.GetAllPlant()
	pclient.GetPlant()

}
