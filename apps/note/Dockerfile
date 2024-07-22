FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY libs/grpc/go.mod libs/grpc/go.sum ./libs/grpc/
COPY apps/note/go.mod apps/note/go.sum ./apps/note/

RUN go work init && \
    go work use ./libs/grpc && \
    go work use ./apps/note && \
    go mod download

COPY libs/grpc ./libs/grpc
COPY apps/note ./apps/note

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/note ./apps/note/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/note/configs/config.yaml ./configs/config.yaml

COPY --from=builder /app/note /app/note

EXPOSE 50051

CMD ["/app/note"]
