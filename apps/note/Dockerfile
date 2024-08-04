FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/note/go.mod apps/note/go.sum /app/apps/note/

RUN go work init && \
    go work use /app/proto && \
    go work use /app/apps/note && \
    go mod download

COPY proto /app/proto
COPY apps/note /app/apps/note

RUN protoc --proto_path=/app/proto/note \
    --go_out=/app/proto/note \
    --go-grpc_out=/app/proto/note \
    /app/proto/note/note.proto

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/note /app/apps/note/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/note/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/note /app/note

EXPOSE 50052

CMD ["/app/note"]
