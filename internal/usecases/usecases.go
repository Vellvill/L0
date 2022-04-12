package usecases

import (
	"L0/internal/model"
	"context"
)

type Repository interface {
	AddModel(ctx context.Context, model *model.Model) error
}
