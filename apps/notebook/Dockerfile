FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/notebook/go.mod apps/notebook/go.sum /app/apps/notebook/

RUN go work init && \
    go work use /app/proto && \
    go work use /app/apps/notebook && \
    go mod download

COPY proto /app/proto
COPY apps/notebook /app/apps/notebook

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /temp/notebook /app/apps/notebook/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/notebook/configs/config.yaml /app/config.yaml
COPY --from=builder /temp/notebook /app/notebook

EXPOSE 50051

CMD ["/app/notebook", "--config", "/app/config.yaml"]
