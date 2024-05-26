package storage

import (
	"fmt"
	"log/slog"
	"sync"
)

type MemoryStorage struct {
	mu   sync.Mutex
	data map[string]string
	logg *slog.Logger
}

func NewMemoryStorage(logg *slog.Logger) *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
		logg: logg,
	}
}

func (s *MemoryStorage) Set(key string, value string) {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
	s.logg.Info(fmt.Sprintf("Set key: %s, value: %s", key, value))
}

func (s *MemoryStorage) Get(key string) (string, bool) {
	s.mu.Lock()
	val, ok := s.data[key]
	s.mu.Unlock()
	return val, ok
}

func (s *MemoryStorage) Delete(key string) {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
	s.logg.Info(fmt.Sprintf("Delete key: %s", key))
}
