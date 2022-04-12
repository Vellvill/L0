package postgres

import (
	"L0/internal/config"
	"L0/internal/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

func NewClient(ctx context.Context, config *config.Config) (pool *pgxpool.Pool, err error) {
	dsn := config.Db.Dsn
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, 5, 5*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	if err = MigrateDatabase(pool, config.Db.MigrationsPath, ctx); err != nil {
		return pool, fmt.Errorf("Unable to migrate, error: %s\n", err)
	}

	log.Println("Connected to postgres database")
	return pool, nil
}
