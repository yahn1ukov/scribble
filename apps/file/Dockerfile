FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY libs/grpc/go.mod libs/grpc/go.sum /app/libs/grpc/
COPY apps/file/go.mod apps/file/go.sum /app/apps/file/

RUN go work init && \
    go work use /app/libs/grpc && \
    go work use /app/apps/file && \
    go mod download

COPY libs/grpc /app/libs/grpc
COPY apps/file /app/apps/file

RUN protoc -I /app/libs/grpc/file \
    --go_out=/app/libs/grpc/file \
    --go-grpc_out=/app/libs/grpc/file \
    /app/libs/grpc/file/file.proto

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/file /app/apps/file/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/file/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/file /app/file

EXPOSE 50052

CMD ["/app/file"]
