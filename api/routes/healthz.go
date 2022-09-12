package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"example.choreo.dev/internal/config"
)

func RegisterHealthRoutes(r fiber.Router) {
	r.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message":   "choreo-example-app is healthy",
			"env":       config.GetConfig().Env,
			"timestamp": time.Now(),
		})
	})
}
