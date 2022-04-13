package reposytories

import (
	"L0/internal/model"
	"L0/internal/usecases"
	"context"
	"encoding/json"
	"fmt"
	pgconn "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repositoryPostgres struct {
	client *pgxpool.Pool
	Hash   *Hash
}

func NewRepository(client *pgxpool.Pool) (usecases.Repository, error) {
	repoDB := &repositoryPostgres{
		client: client,
		Hash:   NewHash(),
	}
	return repoDB, nil
}

func (r *repositoryPostgres) AddModel(ctx context.Context, model *model.Model, uuid string) error {
	q := `
	INSERT INTO models
	(id, model)
	VALUES 
	($1, $2)
	returning id 
`

	send, err := json.Marshal(model.Json)
	if err != nil {
		return err
	}

	if err := r.client.QueryRow(context.Background(), q, model.OrderUID, send).Scan(&model.OrderUID); err != nil {
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

func (r *repositoryPostgres) UpdateHash(ctx context.Context) (err error) {

	q := `
	SELECT id, model
	FROM models
`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return err
	}

	array := make([]model.Model, 0)

	for rows.Next() {
		var models model.Model

		err = rows.Scan(&models.OrderUID, &models.Json)
		if err != nil {
			return err
		}

		array = append(array, models)
	}

	r.Hash.UpdateHash(array)
	return nil
}

func (r *repositoryPostgres) FindInHash(uuid string) ([]byte, error) {
	js, err := r.Hash.FindById(uuid)
	if err != nil {
		return nil, err
	}
	return js, nil
}
