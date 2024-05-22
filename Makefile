CLIENT_ARGS=""
SERVER_ARGS=""

server: build
#	@go run cmd/server/main.go $(SERVER_ARGS)
	@bin/server $(SERVER_ARGS)

client: build
#	@go run cmd/client/main.go $(CLIENT_ARGS)
	@bin/client $(CLIENT_ARGS)

build:
	@go build -o bin/client cmd/client/main.go
	@go build -o bin/server cmd/server/main.go
