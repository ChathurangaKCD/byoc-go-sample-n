package models

import (
	"context"
)

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	Add(ctx context.Context, category Category)
}
