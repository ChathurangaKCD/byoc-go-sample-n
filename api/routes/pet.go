package routes

import (
	"github.com/gofiber/fiber/v2"

	"example.choreo.dev/internal/controllers"
	"example.choreo.dev/internal/utils"
)

func RegisterPetRoutes(router fiber.Router) {
	r := router.Group("/pet")
	r.Post("/", CreatePet)
	r.Get("/:id", GetPet)
}

// CreatePet
// @Param id path int true "Pet ID"
// @Router /api/v1/pet/{id} [post]
// @Success 200 {object} controllers.AddPetRequest "ok"
func CreatePet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	var req controllers.AddPetRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"details": "failed to parse payload",
		})
	}

	res, err := petController.AddPet(ctx, req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetPet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	id := c.Params("id")

	res, err := petController.GetPetById(ctx, id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
