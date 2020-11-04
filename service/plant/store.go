package plant

import (
	"sync"

	"github.com/saravase/golang_grpc_jwt_swagger/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlantStore interface {
	FindAll()
	Find(id *pb.PlantID)
	Save(plant *pb.Plant)
	Update(plant *pb.Plant)
	Delete(id *pb.PlantID)
}

type InMemoryPlantStore struct {
	mutex    sync.RWMutex
	plantMap map[string]*pb.Plant
}

func NewInMemoryPlantStore() *InMemoryPlantStore {
	return &InMemoryPlantStore{
		plantMap: make(map[string]*pb.Plant),
	}
}

// FindAll function get all the plants data from the plantstore
func (store *InMemoryPlantStore) FindAll() (map[string]*pb.Plant, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if len(store.plantMap) == 0 {
		return nil, status.Error(codes.NotFound, "No plant records")
	}
	return store.plantMap, nil
}

// Find function get the plant data based on plant id into the plantstore
func (store *InMemoryPlantStore) Find(plantId *pb.PlantID) (*pb.Plant, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	id := plantId.Value
	if val, found := store.plantMap[id]; found {
		return val, nil
	}
	return nil, status.Errorf(codes.NotFound, "Plant record id %s not found", id)
}

// Save function insert new plant data into the plantstore
func (store *InMemoryPlantStore) Save(plant *pb.Plant) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.plantMap[plant.Id] = plant
}

// Update function update the plant data based on plant id into the plantstore
func (store *InMemoryPlantStore) Update(plant *pb.Plant) (string, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	id := plant.Id
	if _, found := store.plantMap[id]; found {
		store.plantMap[id] = plant
		return id, nil
	}
	return "", status.Errorf(codes.NotFound, "Plant record id %s not found", id)
}

// Delete function delete the plant data based on plant id into the plantstore
func (store *InMemoryPlantStore) Delete(plantId *pb.PlantID) (string, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	id := plantId.Value
	if _, found := store.plantMap[id]; found {
		delete(store.plantMap, id)
		return id, nil
	}
	return "", status.Errorf(codes.NotFound, "Plant record id %s not found", id)
}
