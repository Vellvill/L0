package service

import (
	"L0/internal/usecases"
)

type Implementation struct {
	repo usecases.Repository
}

func New(repo usecases.Repository) (Implementation, error) {
	return Implementation{repo: repo}, nil
}
