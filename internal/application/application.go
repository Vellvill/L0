package application

import (
	"L0/internal/config"
	"L0/internal/service"
	"L0/internal/usecases"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type Application struct {
	config     *config.Config
	router     *mux.Router
	connection *pgxpool.Pool
	repository *usecases.Repository
}

func NewApplication(config *config.Config, pool *pgxpool.Pool, repo *usecases.Repository) (*Application, error) {
	a := &Application{
		config:     config,
		router:     NewRouter(),
		connection: pool,
		repository: service.New(repo),
	}
	return a, nil
}

func (a *Application) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", a.config.Server.Port), a.router)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", nil)
	return r
}
