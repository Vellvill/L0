package reposytories

import (
	"L0/internal/model"
	"L0/internal/usecases"
	"context"
	"sync"
)

type repositoryHash struct {
	sync.Mutex
	hash map[int]model.Model
}

func NewHashRepository() (usecases.Repository, error) {
	repo := &repositoryHash{
		hash: make(map[int]model.Model),
	}
	return repo, nil
}

func (r *repositoryHash) AddModel(ctx context.Context, model *model.Model, uuid string) error {
	return nil
}
