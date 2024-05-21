package compute

import "client-server-db/internal/storage"

func HandlerMessages(req string, s storage.Storage) string {
	command, args, err := Parse(req)
	if err != nil {
		return err.Error()
	}

	query, err := Analyze(command, args, s)
	if err != nil {
		return err.Error()
	}

	return query
}
