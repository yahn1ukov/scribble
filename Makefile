.PHONY: compose-up migrate-create migrate-up generate-proto

compose-up:
	docker-compose -f $(FILE) up -d

migrate-create:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate-up:
	migrate -database=$(URL) -path=./migrations up

proto-generate:
	protoc --go_out=./ --go_opt=paths=source_relative \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	./proto/$(FILE)
