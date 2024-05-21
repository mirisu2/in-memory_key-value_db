package compute

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockStorage is a mock type for the Storage interface
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Set(key, value string) {
	m.Called(key, value)
}

func (m *MockStorage) Get(key string) (string, bool) {
	args := m.Called(key)
	return args.String(0), args.Bool(1)
}

func (m *MockStorage) Delete(key string) {
	m.Called(key)
}

func TestAnalyzeSet(t *testing.T) {
	mockStorage := new(MockStorage)
	mockStorage.On("Set", "key", "value").Return()

	response, err := Analyze("SET", []string{"key", "value"}, mockStorage)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if response != "OK" {
		t.Errorf("Expected response 'OK', got: %s", response)
	}

	mockStorage.AssertExpectations(t)
}

func TestAnalyzeGet(t *testing.T) {
	mockStorage := new(MockStorage)
	mockStorage.On("Get", "key").Return("value", true)

	response, err := Analyze("GET", []string{"key"}, mockStorage)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if response != "value" {
		t.Errorf("Expected response 'value', got: %s", response)
	}

	mockStorage.AssertExpectations(t)
}

func TestAnalyzeDelete(t *testing.T) {
	mockStorage := new(MockStorage)
	mockStorage.On("Delete", "key").Return()

	response, err := Analyze("DELETE", []string{"key"}, mockStorage)
	if err != nil {
		t.Errorf("Error should be nil, got: %v", err)
	}
	if response != "OK" {
		t.Errorf("Expected response 'OK', got: %s", response)
	}

	mockStorage.AssertExpectations(t)
}

func TestAnalyzeUnknownCommand(t *testing.T) {
	mockStorage := new(MockStorage)

	response, err := Analyze("UNKNOWN", []string{"key"}, mockStorage)
	if err == nil || err.Error() != "unknown command" {
		t.Errorf("Expected 'unknown command' error, got: %v", err)
	}
	if response != "" {
		t.Errorf("Expected empty response, got: %s", response)
	}
}

func TestAnalyzeInvalidArguments(t *testing.T) {
	mockStorage := new(MockStorage)

	_, err := Analyze("SET", []string{"key"}, mockStorage)
	if err == nil || err.Error() != "SET requires exactly two arguments" {
		t.Errorf("Expected 'SET requires exactly two arguments' error, got: %v", err)
	}

	_, err = Analyze("GET", []string{}, mockStorage)
	if err == nil || err.Error() != "GET requires exactly one argument" {
		t.Errorf("Expected 'GET requires exactly one argument' error, got: %v", err)
	}

	_, err = Analyze("DELETE", []string{"key1", "key2"}, mockStorage)
	if err == nil || err.Error() != "DELETE requires exactly one argument" {
		t.Errorf("Expected 'DELETE requires exactly one argument' error, got: %v", err)
	}
}
