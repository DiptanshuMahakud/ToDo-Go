# -------- Build stage --------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o todo-api ./cmd/api

# -------- Runtime stage --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/todo-api .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/todo-api"]