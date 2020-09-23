BIN_NAME=clerk-cli

.PHONY: test
test:
	go test -v pkg/actions/*.go

.PHONY:
bin:
	go build -o $(BIN_NAME) cmd/clerk/*.go
