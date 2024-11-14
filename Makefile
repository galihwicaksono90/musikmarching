DB_URL=postgres://admin:root@localhost:5432/musikmarching-db?sslmode=disable
DB_DIR=./db/migration

.PHONY: migration-new 
migration-new:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" create "$(title)" sql

.PHONY: migration-reset 
migration-reset:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" reset

.PHONY: migration-up
migration-up:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" down

.PHONY: migration-status
migration-status:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" status

.PHONY: dev
dev:
	air

.PHONY: tailwind-watch
tailwind-watch:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-watch
templ-watch:
	templ generate -watch

.PHONY: templ
templ:
	templ generate

.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate
