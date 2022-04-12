package reposytories

import (
	"L0/internal/model"
	"L0/internal/usecases"
	"context"
	"encoding/json"
	"fmt"
	pgconn "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"sync"
)

type repositoryPostgres struct {
	client *pgxpool.Pool
	Hash   Hash
}

type Hash struct {
	sync.Mutex
	hash map[string][]byte
}

var Once sync.Once

func NewPostgreshRepository(client *pgxpool.Pool) (usecases.Repository, error) {
	repo := &repositoryPostgres{
		client: client,
		Hash: struct {
			sync.Mutex
			hash map[string][]byte
		}{hash: make(map[string][]byte)},
	}
	Once.Do(func() {
		err := repo.UpdateHash(context.Background())
		if err != nil {
			log.Println(err)
		}
	})

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

	if err := r.client.QueryRow(context.Background(), q, model.OrderUID, model.Json).Scan(); err != nil {
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

	for rows.Next() {
		var models model.Model

		err = rows.Scan(&models.OrderUID, &models.Json)
		if err != nil {
			return err
		}

		go func() {
			err := r.Hash.UpdateHash()
			if err != nil {
				log.Println(err)
			}
		}()
	}
	return nil
}

func (h *Hash) AddModelHash(model model.Model, uuid string) (err error) {
	if _, ok := h.hash[uuid]; ok {
		return nil
	}
	h.Lock()
	h.hash[uuid], err = json.Marshal(model.Json)
	if err != nil {
		return err
	}
	h.Unlock()
	return nil
}

func (h *Hash) UpdateHash(models ...model.Model) (err error) {
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(len(models))
		go func() {
			for _, v := range models {
				h.Lock()
				h.hash[v.OrderUID], err = json.Marshal(v.Json)
				if err != nil {
					log.Println(err)
				}
				h.Unlock()
				wg.Done()
			}
		}()
		wg.Wait()
	}()
	return nil
}
