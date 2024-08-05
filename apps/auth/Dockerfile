FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/hash/go.mod libs/hash/go.sum /app/libs/hash/
COPY libs/jwt/go.mod libs/jwt/go.sum /app/libs/jwt/
COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/auth/go.mod apps/auth/go.sum /app/apps/auth/

RUN go work init && \
    go work use /app/libs/hash && \
    go work use /app/libs/jwt && \
    go work use /app/proto && \
    go work use /app/apps/auth && \
    go mod download

COPY libs/hash /app/libs/hash
COPY libs/jwt /app/libs/jwt
COPY proto /app/proto
COPY apps/auth /app/apps/auth

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/auth /app/apps/auth/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/auth/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/auth /app/auth

EXPOSE 50054

CMD ["/app/auth"]
