FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/ ./

COPY proto/ ./proto/

ENV CGO_ENABLED=1
ENV GOOS=linux
RUN go build -a -ldflags '-extldflags "-static"' -o server cmd/server/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates sqlite

COPY --from=builder /app/server /app/server

RUN mkdir -p /app/data

EXPOSE ${GO_PORT}
EXPOSE ${GRPC_PORT}

RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]