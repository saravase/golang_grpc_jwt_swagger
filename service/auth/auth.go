package auth

import (
	"fmt"
	"time"

	"github.com/saravase/golang_grpc_jwt_swagger/service/user"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	secretKey  string
	tkDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json: "username"`
	Role     string `json: "role"`
}

func NewAuth(secretKey string, tkDuration time.Duration) *Auth {
	return &Auth{
		secretKey:  secretKey,
		tkDuration: tkDuration,
	}
}

func (auth *Auth) Generate(user *user.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth.tkDuration).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(auth.secretKey))

}

func (auth *Auth) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected token sigining method")
			}

			return []byte(auth.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Invalid access token: %v", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil

}
