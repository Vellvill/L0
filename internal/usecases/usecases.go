package usecases

import (
	"L0/internal/model"
	"context"
)

type Repository interface {
	AddModel(ctx context.Context, model *model.Model, uuid string) error
	UpdateHash(ctx context.Context) (err error)
	FindInHash(uuid string) ([]byte, error)
}
