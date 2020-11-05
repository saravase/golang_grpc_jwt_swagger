package main

import (
	"context"
	"log"
	"time"

	"github.com/saravase/golang_grpc_jwt_swagger/client/auth"
	"github.com/saravase/golang_grpc_jwt_swagger/client/plant"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:9090"
	refreshDuration = time.Minute
)

func getAuthMethods() map[string]bool {

	plantServicePath := "/pb.AuthService/"

	return map[string]bool{
		plantServicePath + "CreatePlant": true,
		plantServicePath + "UpdatePlant": true,
		plantServicePath + "DeletePlant": true,
		plantServicePath + "GetPlants":   true,
		plantServicePath + "GetPlant":    true,
	}
}

func main() {

	conn1, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("% connection creataion failed", address)
	}
	defer conn1.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	aclient := auth.NewAuthClient(conn1, ctx)
	aclient.Register()
	interceptor, err := auth.NewAuthInterceptor(aclient, getAuthMethods(), refreshDuration)
	if err != nil {
		log.Fatalf("Auth interceptor creation failed. Reason: %v", err)
	}

	conn2, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatalf("% connection creataion failed", address)
	}
	defer conn2.Close()

	pclient := plant.NewPlantClient(conn2, ctx)
	pclient.CreatePlant()
	pclient.GetPlants()
	pclient.UpdatePlant()
	pclient.GetPlants()
	pclient.DeletePlant()
	pclient.GetPlants()
	pclient.GetPlant()

}
