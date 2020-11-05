package auth

import (
	"log"

	"github.com/saravase/golang_grpc_jwt_swagger/service/user"
)

type AuthServer struct {
	userStore user.UserStore
	auth      *Auth
	logger    *log.Logger
}

func NewAuthServer(store user.UserStore, auth *Auth, logger *log.Logger) *AuthServer {
	return &AuthServer{
		userStore: store,
		auth:      auth,
		logger:    logger,
	}
}
