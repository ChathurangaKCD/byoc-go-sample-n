package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

const correlationIdKeyName = "x-correlation-id"

func GetRequestContext(rCtx *fiber.Ctx) context.Context {
	correlationId := rCtx.Get(correlationIdKeyName)
	ctx := context.Background()
	return context.WithValue(ctx, correlationIdKeyName, correlationId)
}
