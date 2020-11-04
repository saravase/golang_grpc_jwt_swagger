package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	secretKey  string
	tkduration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json: "username"`
	Role     string `json: "role"`
}
