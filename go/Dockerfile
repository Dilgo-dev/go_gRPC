FROM golang:1.23-alpine AS development

RUN apk add --no-cache git bash curl gcc musl-dev sqlite-dev && \
    go install github.com/air-verse/air@latest

ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3333

ENV GO_ENV=development
ENV PORT=3333

CMD ["air", "-c", ".air.toml"]
