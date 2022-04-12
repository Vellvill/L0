package reposytories

import (
	"L0/internal/model"
	"L0/internal/usecases"
	"context"
	"fmt"
	pgconn "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repositoryPostgres struct {
	pool *pgxpool.Pool
}

func NewPostgreshRepository(pool *pgxpool.Pool) (usecases.Repository, error) {
	repo := &repositoryPostgres{
		pool: pool,
	}
	return repo, nil
}

func (r *repositoryPostgres) AddModel(ctx context.Context, model *model.Model, uuid string) error {

	q := `
	INSERT INTO models
	(id, model)
	VALUES 
	($1, $2)
	returning id 
`

	if err := r.pool.QueryRow(context.Background(), q, model.OrderUID, model.Json).Scan(); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf(
				"SQL Error: %s, Detail: %s, Where: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState()))
			return newErr
		} else {
			return err
		}
	}
	return nil
}
