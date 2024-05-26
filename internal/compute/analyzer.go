package compute

import (
	"client-server-db/internal/storage"
	"errors"
)

var (
	ErrSet            = errors.New("SET requires exactly two arguments")
	ErrGet            = errors.New("GET requires exactly one argument")
	ErrKeyNotFound    = errors.New("key not found")
	ErrDelete         = errors.New("DELETE requires exactly one argument")
	ErrUnknownCommand = errors.New("unknown command")
)

func Analyze(command string, args []string, storage storage.Storage) (response string, err error) {
	switch command {
	case "SET":
		if len(args) != 2 {
			return "", ErrSet
		}
		key := args[0]
		value := args[1]
		storage.Set(key, value)
		return "OK", nil
	case "GET":
		if len(args) != 1 {
			return "", ErrGet
		}
		key := args[0]
		value, ok := storage.Get(key)
		if !ok {
			return "", ErrKeyNotFound
		}
		return value, nil
	case "DELETE":
		if len(args) != 1 {
			return "", ErrDelete
		}
		key := args[0]
		storage.Delete(key)
		return "OK", nil
	default:
		return "", ErrUnknownCommand
	}
}
