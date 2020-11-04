package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string
	EncryptPassword string
	Role            string
}

func NewUser(
	username string,
	password string,
	role string) (*User, error) {

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Password encryption failed. Reason: %v", err)
	}

	user := &User{
		Username:        username,
		EncryptPassword: string(encryptPassword),
		Role:            role,
	}
	return user, nil
}

func (user *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptPassword), []byte(password))
	return err == nil
}

func (user *User) Clone() *User {
	return &User{
		Username:        user.Username,
		EncryptPassword: user.EncryptPassword,
		Role:            user.Role,
	}
}
