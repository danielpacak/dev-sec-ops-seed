SOURCES := $(shell find . -name '*.go')
BINARY := seed

build: $(BINARY)

$(BINARY): $(SOURCES)
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -o bin/$(BINARY) cmd/seed/main.go

test: build
	GO111MODULE=on go test -v -short -race -coverprofile=coverage.txt -covermode=atomic ./...