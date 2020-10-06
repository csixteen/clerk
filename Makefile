BIN_NAME=clerk-cli
SERVER_NAME=clerk-server

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY:
bin:
	go build -o $(BIN_NAME) cmd/clerk/*.go

server:
	go build -o $(SERVER_NAME) app/clerk-api/*.go
