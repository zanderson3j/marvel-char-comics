FROM golang:1.16

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/web

EXPOSE 8000
ENTRYPOINT ["/app/web"]