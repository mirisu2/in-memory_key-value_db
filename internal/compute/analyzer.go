package compute

import (
	"client-server-db/internal/storage"
	"errors"
)

func Analyze(command string, args []string, storage storage.Storage) (response string, err error) {
	switch command {
	case "SET":
		if len(args) != 2 {
			return "", errors.New("SET requires exactly two arguments")
		}
		key := args[0]
		value := args[1]
		storage.Set(key, value)
		return "OK", nil
	case "GET":
		if len(args) != 1 {
			return "", errors.New("GET requires exactly one argument")
		}
		key := args[0]
		value, ok := storage.Get(key)
		if !ok {
			return "", errors.New("error: key not found")
		}
		return value, nil
	case "DELETE":
		if len(args) != 1 {
			return "", errors.New("DELETE requires exactly one argument")
		}
		key := args[0]
		storage.Delete(key)
		return "OK", nil
	default:
		return "", errors.New("unknown command")
	}
}
