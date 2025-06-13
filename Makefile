include .env

DB_URL ?= postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR = internal/adapter/storage/postgres/migrations

.PHONY: run lint test build dev swag migration migrateup migrateup1 migratedown migratedown1 docker-migrateup

run:
	go run cmd/main.go

migration:
	@if [ -z "$(name)" ]; then \
		echo "❌ Usage: make migration name=alter_table_xyz"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(name)

migrateup:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose down 1
force:
	@if [ -z "$(version)" ]; then \
		echo "❌ Usage: make force version=1"; \
		exit 1; \
	fi
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(version)

docker-migrateup:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations migrate/migrate -path=/migrations -database="postgres://postgres:postgres@host.docker.internal:5432/aok_connect_individual?sslmode=disable" up
