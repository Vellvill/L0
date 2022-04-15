package reposytories

import (
	"L0/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type Hash struct {
	mu   sync.Mutex
	rw   sync.RWMutex
	Hash map[string][]byte
}

func NewHash() *Hash {
	return &Hash{
		Hash: make(map[string][]byte),
	}
}

func (h *Hash) AddModelHash(model model.Model) (err error) {
	if _, ok := h.Hash[model.OrderUID]; ok {
		return nil
	}
	h.mu.Lock()
	h.Hash[model.OrderUID], err = json.Marshal(model.Json)
	if err != nil {
		return err
	}
	h.mu.Unlock()
	return nil
}

func (h *Hash) UpdateHash(models []model.Model) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(models))
	for _, v := range models {
		go func(v model.Model) {
			defer wg.Done()
			h.mu.Lock()
			h.Hash[v.OrderUID], err = json.Marshal(v.Json)
			if err != nil {
				log.Println(err)
			}
			h.mu.Unlock()
		}(v)
	}
	wg.Wait()
	log.Printf("Data successfully loaded into the cache, data len = %d\n", len(models))
	return nil
}

func (h *Hash) FindById(uuid string) ([]byte, error) {
	h.rw.RLock()
	if j, ok := h.Hash[uuid]; ok {
		h.rw.RUnlock()
		return j, nil
	} else {
		h.rw.RUnlock()
		return nil, fmt.Errorf("Didn't find any jsons with '%s' id\n", uuid)
	}
}
