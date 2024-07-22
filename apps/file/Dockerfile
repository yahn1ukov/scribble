FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/grpc/go.mod libs/grpc/go.sum ./libs/grpc/
COPY apps/file/go.mod apps/file/go.sum ./apps/file/

RUN go work init && \
    go work use ./libs/grpc && \
    go work use ./apps/file && \
    go mod download

COPY libs/grpc ./libs/grpc
COPY apps/file ./apps/file

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/file ./apps/file/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/file/configs/config.yaml ./configs/config.yaml

COPY --from=builder /app/file /app/file

EXPOSE 50052

CMD ["/app/file"]
