package user

import (
	"sync"

	"github.com/saravase/golang_grpc_jwt_swagger/service/constant"
)

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

type InMemoryUserStore struct {
	mutex   sync.RWMutex
	userMap map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		userMap: make(map[string]*User),
	}
}

func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	store.mutex.Unlock()

	if _, found := store.userMap[user.Username]; found {
		return constant.ErrAlreadyExists
	}

	store.userMap[user.Username] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	store.mutex.RUnlock()

	user, found := store.userMap[username]
	if found {
		return user.Clone(), nil
	}

	return nil, constant.ErrNotFound
}
