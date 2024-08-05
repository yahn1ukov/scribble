FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/grpc/go.mod libs/grpc/go.sum /app/libs/grpc/
COPY libs/jwt/go.mod libs/jwt/go.sum /app/libs/jwt/
COPY libs/respond/go.mod /app/libs/respond/
COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/gateway/go.mod apps/gateway/go.sum /app/apps/gateway/

RUN go work init && \
    go work use /app/libs/grpc && \
    go work use /app/libs/jwt && \
    go work use /app/libs/respond && \
    go work use /app/proto && \
    go work use /app/apps/gateway && \
    go mod download

COPY libs/grpc /app/libs/grpc
COPY libs/jwt /app/libs/jwt
COPY libs/respond /app/libs/respond
COPY proto /app/proto
COPY apps/gateway /app/apps/gateway

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/gateway /app/apps/gateway/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/gateway/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/gateway /app/gateway

EXPOSE 8000

CMD ["/app/gateway"]
