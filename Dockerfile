FROM golang:latest

WORKDIR /server

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build initial_router_config.go main.go" --command=./main