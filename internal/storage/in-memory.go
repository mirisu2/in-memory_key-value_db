package storage

import (
	"client-server-db/internal/logger"
	"fmt"
)

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (s *MemoryStorage) Set(key string, value string) {
	s.data[key] = value
	logger.Log.Info(fmt.Sprintf("Set key: %s, value: %s", key, value))
}

func (s *MemoryStorage) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *MemoryStorage) Delete(key string) {
	delete(s.data, key)
	logger.Log.Info(fmt.Sprintf("Delete key: %s", key))
}
