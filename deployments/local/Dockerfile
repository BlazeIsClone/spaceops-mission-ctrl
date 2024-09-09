FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD ../../ /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="make build" --command=./app