# .PHONY is a special target that tells make that the targets listed are not actual files, but rather just labels for commands to be executed.
.PHONY: migrate-up migrate-down migrate-create migrate-drop run

DB_ADDR=postgres://postgres:password@localhost:5432/auth_service_db?sslmode=disable

create-migrate:
	@if [ -z "$(name)" ]; then echo "name is required"; exit 1; fi
	migrate create -ext sql -dir cmd/migrate/migrations -seq $(name)
	# Example: make create-migrate name=create_login_logs

migrate-up:
	migrate -path cmd/migrate/migrations -database "$(DB_ADDR)" -verbose up

migrate-down:
	migrate -path cmd/migrate/migrations -database "$(DB_ADDR)" -verbose down

run:
	go run cmd/api/main.go
