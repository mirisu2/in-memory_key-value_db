package storage

import (
	"testing"
)

func TestNewMemoryStorage(t *testing.T) {
	storage := NewMemoryStorage()
	if storage == nil {
		t.Error("NewMemoryStorage returned nil")
	}
	if len(storage.data) != 0 {
		t.Errorf("NewMemoryStorage data map should be empty, got %v", storage.data)
	}
}

func TestSetAndGet(t *testing.T) {
	storage := NewMemoryStorage()
	key := "testKey"
	value := "testValue"

	storage.Set(key, value)
	retrievedValue, exists := storage.Get(key)
	if !exists {
		t.Errorf("Expected key '%s' to exist", key)
	}
	if retrievedValue != value {
		t.Errorf("Expected value '%s', got '%s'", value, retrievedValue)
	}
}

func TestDelete(t *testing.T) {
	storage := NewMemoryStorage()
	key := "testKey"
	value := "testValue"

	storage.Set(key, value)
	storage.Delete(key)

	_, exists := storage.Get(key)
	if exists {
		t.Errorf("Key '%s' should have been deleted", key)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	storage := NewMemoryStorage()
	key := "nonExistentKey"

	_, exists := storage.Get(key)
	if exists {
		t.Errorf("Expected non-existent key '%s' to not exist", key)
	}
}
