package routes

import (
	"example.choreo.dev/internal/controllers"
	"example.choreo.dev/internal/repositories"
)

var petController *controllers.PetController
var categoryController *controllers.CategoryController

func initControllers() {
	categoryRepository := repositories.NewCategoryRepository()
	petRepository := repositories.NewPetRepository()
	petController = controllers.NewPetController(petRepository, categoryRepository)
	categoryController = controllers.NewCategoryController(categoryRepository)
}
