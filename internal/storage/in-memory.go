package storage

import (
	"client-server-db/internal/logger"
	"fmt"
	"sync"
)

type MemoryStorage struct {
	mu   sync.Mutex
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (s *MemoryStorage) Set(key string, value string) {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
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
