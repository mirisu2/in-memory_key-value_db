package storage

import (
	"errors"
	"fmt"
	"log/slog"
)

type Storage interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Delete(key string)
}

func NewStorage(storageType string, logg *slog.Logger) (Storage, error) {
	switch storageType {
	case "memory":
		return NewMemoryStorage(logg), nil
	default:
		return nil, errors.New(fmt.Sprintf("unknown storage type: %s", storageType))
	}
}
