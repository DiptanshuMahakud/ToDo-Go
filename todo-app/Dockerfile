From golang:1.22-alpine as builder

WORKDIR /app

#cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

#COPY source
COPY . .

#Build binary

Run CGO_ENABLED=0 GOOS=linux GOARCH=amd64\
    go build -o todo-api ./cmd/api
# Runtime Stage

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/todo-api .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/todo-api"]
