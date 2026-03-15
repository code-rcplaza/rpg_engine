# ---- Dev stage (hot reload con Air) ----
FROM golang:1.23-alpine AS dev

RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# ---- Build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /bin/server ./cmd/server

# ---- Production stage ----
FROM alpine:3.20 AS prod

RUN apk add --no-cache ca-certificates sqlite
COPY --from=builder /bin/server /bin/server

EXPOSE 8080
ENTRYPOINT ["/bin/server"]