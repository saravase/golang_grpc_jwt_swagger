package main

import (
	"log"
	"math/rand"
)

type Data struct {
	value int
}

func getAccessibleRoles() map[string][]string {

	const plantServicePath = "/pb.AuthService/"

	return map[string][]string{
		plantServicePath + "CreatePlant": {"admin"},
		plantServicePath + "UpdatePlant": {"admin"},
		plantServicePath + "DeletePlant": {"admin"},
		plantServicePath + "GetPlants":   {"admin", "user"},
		plantServicePath + "GetPlant":    {"admin", "user"},
	}
}

func main() {

	roles := getAccessibleRoles()
	value, ok := roles["/pb.AuthService/DeletePlant"]
	log.Print(value)
	log.Print(ok)
}

func (data *Data) RefershData() {
	data.value = rand.Intn(100)
	log.Printf("Data updated: %v", data)
}
