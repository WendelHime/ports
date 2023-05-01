generate:
	go generate github.com/WendelHime/ports/...

build:
	docker-compose build

run:
	docker-compose up

test:
	go test -race github.com/WendelHime/ports/...

lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v

all: generate test build run
