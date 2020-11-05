package auth

import (
	"fmt"
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
)

func (client *AuthClient) Login() (string, error) {

	user := &pb.LoginRequest{
		Username: "optimus",
		Password: "optimus",
	}

	token, err := client.cli.Login(client.ctx, user)

	if err != nil {
		log.Fatalf("[ERROR] While Login user %v", err)
		return "", fmt.Errorf("Login failed. Reason : %v", err)
	}
	log.Printf("Logged user access token %s", token)

	return token.AccessToken, nil
}
