CLIENT_ARGS=""
SERVER_ARGS=""

server:
	@go run cmd/server/main.go $(SERVER_ARGS)

client:
	@go run cmd/client/main.go $(CLIENT_ARGS)

build:
	@go build -o bin/client cmd/client/main.go
	@go build -o bin/server cmd/server/main.go