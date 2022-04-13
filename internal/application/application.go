package application

import (
	"L0/internal/config"
	"L0/internal/model"
	"L0/internal/usecases"
	"L0/internal/utils"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"strings"
	"time"
)

type Application struct {
	config     *config.Config
	router     *mux.Router
	connection *pgxpool.Pool
	repository usecases.Repository
}

func NewApplication(config *config.Config, pool *pgxpool.Pool, repDB usecases.Repository) (*Application, error) {
	a := &Application{
		config:     config,
		router:     nil,
		connection: pool,
		repository: repDB,
	}

	a.NewRouter()

	return a, nil
}

func (a *Application) Start(ctx context.Context) error {

	err := a.repository.UpdateHash(context.Background())

	if err != nil {
		log.Println(err)
	}

	var sc stan.Conn

	err = utils.DoWithTries(func() error {
		sc, err = stan.Connect(a.config.Nats.ClusterID, a.config.Nats.ClientID)
		if err != nil {
			return err
		}
		return nil
	}, 5, 5*time.Second)

	update := make(chan *model.Model)

	var _, er = sc.Subscribe(a.config.Nats.Channel, func(m *stan.Msg) {
		js, err := model.NewModel(m.Data)
		if err != nil {
			return
		}
		update <- js
	})
	if er != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case x := <-update:
				err := a.repository.AddModel(ctx, x, x.OrderUID)
				if err != nil {
					log.Println(err)
				}
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return http.ListenAndServe(fmt.Sprintf(":%s", a.config.Server.Port), a.router)
}

func (a *Application) NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", a.FindById)
	a.router = r
}

func (a *Application) FindById(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Trim(r.URL.Path, "/")
	data, err := a.repository.FindInHash(uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
