# dockerfile
FROM golang:1.22-alpine AS builder

COPY . /github.com/mchekalov/auth/source
WORKDIR /github.com/mchekalov/auth/source

RUN go mod download
RUN go build -o ./bin/crud_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/mchekalov/auth/source/bin/crud_server .
ADD .env .

CMD ["./crud_server"]