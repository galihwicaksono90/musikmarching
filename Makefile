include .env

DB_URL="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

.PHONY: migration-new 
migration-new:
	@goose -dir "$(DB_DIR)" postgres "$(DB_URL)" create "$(title)" sql

.PHONY: migration-reset 
migration-reset:
	@goose -dir "$(DB_DIR)" postgres "$(DB_URL)" reset

.PHONY: migration-up
migration-up:
	@goose -dir "$(DB_DIR)" postgres "$(DB_URL)" up

.PHONY: migration-down
migration-down:
	@goose -dir "$(DB_DIR)" postgres "$(DB_URL)" down

.PHONY: migration-status
migration-status:
	@goose -dir "$(DB_DIR)" postgres "$(DB_URL)" status

.PHONY: migration-seed
migration-seed:
	@goose -dir "$(DB_SEED_DIR)" -no-versioning  postgres "$(DB_URL)" up

.PHONY: sqlc-generate
sqlc-generate:
	@sqlc generate

.PHONY: dev
dev:
	@air
