FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY libs/grpc/go.mod libs/grpc/go.sum /app/libs/grpc/
COPY apps/note/go.mod apps/note/go.sum /app/apps/note/

RUN go work init && \
    go work use /app/libs/grpc && \
    go work use /app/apps/note && \
    go mod download

COPY libs/grpc /app/libs/grpc
COPY apps/note /app/apps/note

RUN protoc -I /app/libs/grpc/note \
    --go_out=/app/libs/grpc/note \
    --go-grpc_out=/app/libs/grpc/note \
    /app/libs/grpc/note/note.proto

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/note /app/apps/note/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/note/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/note /app/note

EXPOSE 50051

CMD ["/app/note"]
