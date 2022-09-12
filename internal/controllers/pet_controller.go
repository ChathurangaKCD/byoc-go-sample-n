package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"

	"example.choreo.dev/internal/models"
	"example.choreo.dev/internal/repositories"
)

type PetController struct {
	petRepository models.PetRepository
}

func NewPetController(petRepository models.PetRepository) *PetController {
	return &PetController{petRepository: petRepository}
}

func (c PetController) AddPet(ctx context.Context, request AddPetRequest) (*models.Pet, error) {
	var pet models.Pet
	if err := copier.Copy(&pet, request); err != nil {
		return nil, err
	}
	err := c.petRepository.Add(ctx, &pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c PetController) GetPetById(ctx context.Context, petId string) (*models.Pet, error) {
	pet, err := c.petRepository.GetPetById(ctx, petId)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return nil, fiber.NewError(http.StatusNotFound, "record not found")
		}
		return nil, err
	}
	return pet, nil
}
