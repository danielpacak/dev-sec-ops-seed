SOURCES := $(shell find . -name '*.go')
BINARY := seed
IMAGE_TAG := dev
IMAGE := docker.pkg.github.com/danielpacak/dev-sec-ops-seed/seed:$(IMAGE_TAG)

build: $(BINARY)

$(BINARY): $(SOURCES)
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -o $(BINARY) cmd/seed/main.go

test: build
	GO111MODULE=on go test -v -short -race -coverprofile=coverage.txt -covermode=atomic ./...

container-build: build
	docker build -t $(IMAGE) .

container-run: container-build
	docker run --rm --name seed -p 8080:8080 $(IMAGE)