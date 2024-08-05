FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/hash/go.mod libs/hash/go.sum /app/libs/hash/
COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/user/go.mod apps/user/go.sum /app/apps/user/

RUN go work init && \
    go work use /app/libs/hash && \
    go work use /app/proto && \
    go work use /app/apps/user && \
    go mod download

COPY libs/hash /app/libs/hash
COPY proto /app/proto
COPY apps/user /app/apps/user

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/user /app/apps/user/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/user/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/user /app/user

EXPOSE 50050

CMD ["/app/user"]
