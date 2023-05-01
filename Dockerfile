FROM golang:1.20-alpine as builder

WORKDIR /build

COPY . .
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver ./cmd/ports-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /build/apiserver /usr/bin/apiserver
EXPOSE 8080
ENTRYPOINT ["/usr/bin/apiserver"]
