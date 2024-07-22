FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/grpc/go.mod libs/grpc/go.sum ./libs/grpc/
COPY libs/respond/go.mod ./libs/respond/
COPY apps/gateway/go.mod apps/gateway/go.sum ./apps/gateway/

RUN go work init && \
    go work use ./libs/grpc && \
    go work use ./libs/respond && \
    go work use ./apps/gateway && \
    go mod download

COPY libs/grpc ./libs/grpc
COPY libs/respond ./libs/respond
COPY apps/gateway ./apps/gateway

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/gateway ./apps/gateway/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/gateway/configs/config.yaml ./configs/config.yaml

COPY --from=builder /app/gateway /app/gateway

EXPOSE 8000

CMD ["/app/gateway"]
