FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/grpc/go.mod libs/grpc/go.sum ./libs/grpc/
COPY apps/notebook/go.mod apps/notebook/go.sum ./apps/notebook/

RUN go work init && \
    go work use ./libs/grpc && \
    go work use ./apps/notebook && \
    go mod download

COPY libs/grpc ./libs/grpc
COPY apps/notebook ./apps/notebook

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/notebook ./apps/notebook/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/notebook/configs/config.yaml ./configs/config.yaml

COPY --from=builder /app/notebook /app/notebook

EXPOSE 50050

CMD ["/app/notebook"]
