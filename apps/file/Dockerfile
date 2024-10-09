FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/file/go.mod apps/file/go.sum /app/apps/file/

RUN go work init && \
    go work use /app/proto && \
    go work use /app/apps/file && \
    go mod download

COPY proto /app/proto
COPY apps/file /app/apps/file

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /temp/file /app/apps/file/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/file/configs/config.yaml /app/config.yaml
COPY --from=builder /temp/file /app/file

EXPOSE 50053

CMD ["/app/file", "--config", "/app/config.yaml"]
