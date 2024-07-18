.PHONY: compose-up migrate-create migrate-up

compose-up:
	docker-compose -f $(FILE) up -d

migrate-create:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate-up:
	migrate -database="$(URL)" -path=./migrations up
