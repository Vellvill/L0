package reposytories

import (
	"L0/internal/config"
	"L0/internal/postgres"
	"L0/internal/usecases"
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL, migrationsPath string) (usecases.Repository, func(...string)) {
	t.Helper()

	config, err := config.GetConfig()
	config.Db.Dsn = databaseURL

	client, err := postgres.NewClient(context.Background(), config)

	repo, err := NewRepository(client)
	if err != nil {
		t.Fatal(err)
	}

	err = postgres.MigrateDatabase(client, migrationsPath, context.Background())
	if err != nil {
		t.Fatal(err)
	}

	return repo, func(tables ...string) {
		if len(tables) > 0 {
			q := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if _, err := client.Exec(context.Background(), q); err != nil {
				t.Fatal(err)
			}
			return
		}
	}
}
