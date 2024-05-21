package storage

import (
	"errors"
	"fmt"
)

type Storage interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Delete(key string)
}

func NewStorage(storageType string) (Storage, error) {
	switch storageType {
	case "memory":
		return NewMemoryStorage(), nil
	default:
		return nil, errors.New(fmt.Sprintf("unknown storage type: %s", storageType))
	}
}
