package plant

import "log"

type PlantServer struct {
	logger *log.Logger
	store  *InMemoryPlantStore
}

func NewPlantServer(logger *log.Logger, store *InMemoryPlantStore) *PlantServer {
	return &PlantServer{
		logger: logger,
		store:  store,
	}
}
