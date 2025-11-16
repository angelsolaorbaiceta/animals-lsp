.PHONY: run
run: # Starts the language server attached to the host terminal
	go run cmd/animals-lsp/main.go

.PHONY: build
build: # Builds the LSP binary into the build/ directory
	go build -o build/animals-lsp cmd/animals-lsp/main.go

.PHONY: test
test: # Runs all tests
	go test -v ./...