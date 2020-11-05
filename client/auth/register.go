package auth

import (
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
)

func (client *AuthClient) Register() {

	users := []*pb.RegisterRequest{
		&pb.RegisterRequest{
			Username: "primz",
			Password: "primz",
			Role:     "admin",
		},
		&pb.RegisterRequest{
			Username: "optimus",
			Password: "optimus",
			Role:     "user",
		},
	}

	for _, user := range users {
		username, err := client.cli.Register(client.ctx, user)

		if err != nil {
			log.Fatalf("[ERROR] While register user %v", err)
		}
		log.Printf("Registered user name %s", username)
	}

}
