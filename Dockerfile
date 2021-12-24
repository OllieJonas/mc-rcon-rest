FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /test-rcon

EXPOSE 8080

ENTRYPOINT [ "/test-rcon --mode release" ]
