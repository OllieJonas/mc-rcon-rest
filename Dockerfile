# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/rest-rcon /app/
EXPOSE 8080
ENTRYPOINT ["/app/rest-rcon"]

# go build -t rest-rcon .
# docker run --publish 8080:8080 rest-rcon [args]
