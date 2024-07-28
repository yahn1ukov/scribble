FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY libs/grpc/go.mod libs/grpc/go.sum /app/libs/grpc/
COPY apps/notebook/go.mod apps/notebook/go.sum /app/apps/notebook/

RUN go work init && \
    go work use /app/libs/grpc && \
    go work use /app/apps/notebook && \
    go mod download

COPY libs/grpc /app/libs/grpc
COPY apps/notebook /app/apps/notebook

RUN protoc -I /app/libs/grpc/notebook \
    --go_out=/app/libs/grpc/notebook \
    --go-grpc_out=/app/libs/grpc/notebook \
    /app/libs/grpc/notebook/notebook.proto

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/notebook /app/apps/notebook/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/notebook/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/notebook /app/notebook

EXPOSE 50050

CMD ["/app/notebook"]
