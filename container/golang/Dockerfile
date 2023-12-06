FROM golang:1.20 as build

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.29.0

COPY ./src/go.mod /app
COPY ./src/go.sum /app
RUN go mod download
COPY ./src /app
RUN CGO_ENABLED=0 go build -o main main.go

