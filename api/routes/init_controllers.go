package routes

import (
	"example.choreo.dev/internal/controllers"
	"example.choreo.dev/internal/repositories"
)

var petController *controllers.PetController

func initControllers() {
	petRepository := repositories.NewPetRepository()
	petController = controllers.NewPetController(petRepository)
}
