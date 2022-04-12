package utils

import (
	"log"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) error {
	for i := attempts; i > 0; i-- {
		if err := fn(); err != nil {
			log.Println("error while doing connection, err:%s\n", err)
			time.Sleep(delay)
		}
		return nil
	}
	return nil
}
