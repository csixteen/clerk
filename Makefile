BIN_NAME=clerk-cli

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY:
bin:
	CGO_ENABLED=0 go build -o $(BIN_NAME) cmd/clerk/*.go
