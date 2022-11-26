FROM golang:1.18.2
WORKDIR /go/src/docker-go-api
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]