package reposytories

import (
	"L0/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type Hash struct {
	sync.Mutex
	Hash map[string][]byte
}

func NewHash() *Hash {
	return &Hash{
		Hash: make(map[string][]byte),
	}
}

func (h *Hash) AddModelHash(model model.Model, uuid string) (err error) {
	if _, ok := h.Hash[uuid]; ok {
		return nil
	}
	h.Lock()
	h.Hash[uuid], err = json.Marshal(model.Json)
	if err != nil {
		return err
	}
	h.Unlock()
	return nil
}

func (h *Hash) UpdateHash(models []model.Model) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(models))
	go func() {
		for _, v := range models {
			h.Lock()
			h.Hash[v.OrderUID], err = json.Marshal(v.Json)
			if err != nil {
				log.Println(err)
			}
			h.Unlock()
			wg.Done()
		}
	}()
	wg.Wait()
	return nil
}

func (h *Hash) FindById(uuid string) ([]byte, error) {
	if j, ok := h.Hash[uuid]; ok {
		return j, nil
	} else {
		return nil, fmt.Errorf("Didn't find any jsons with '%s' id\n", uuid)
	}
}
