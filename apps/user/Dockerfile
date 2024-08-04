FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/user/go.mod apps/user/go.sum /app/apps/user/

RUN go work init && \
    go work use /app/proto && \
    go work use /app/apps/user && \
    go mod download

COPY proto /app/proto
COPY apps/user /app/apps/user

RUN protoc --proto_path=/app/proto/user \
    --go_out=/app/proto/user \
    --go-grpc_out=/app/proto/user \
    /app/proto/user/user.proto

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/user /app/apps/user/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/user/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/user /app/user

EXPOSE 50050

CMD ["/app/user"]
