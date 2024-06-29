.PHONY: compose-up migrate-create migrate-up migrate-down generate-proto

compose-up:
	docker-compose -f ./deployments/$(file) up -d

migrate-create:
	migrate create -ext sql -dir ./scripts/migrations -seq $(name)

migrate-up:
	migrate -database $(url) -path ./scripts/migrations up

migrate-down:
	migrate -database $(url) -path ./scripts/migrations down

generate-proto:
	protoc --go_out=./ --go_opt=paths=source_relative \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	$(file)
