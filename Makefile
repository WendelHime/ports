generate:
	go generate github.com/WendelHime/ports/...

build:
	docker build -t ports:latest .

run:
	docker run --rm -it -p 8080:8080 ports:latest

test:
	go test -race github.com/WendelHime/ports/...

lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v

all: generate test lint build run
