DB_HOST?=localhost
DB_PORT?=5454
DB_USER?=postgres
DB_PASSWORD?=1
DB_NAME?=cinemax
DB_SSLMODE?=disable

MIGRATION_DIR=./db/migrations

MIGRATE=migrate -source "file://$(MIGRATION_DIR)" \
		-database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)"

migration_create:
	migrate create -seq -dir $(MIGRATION_DIR) -ext sql $(name)

migration_up: 
		$(MIGRATE) up

migration_down: 
		$(MIGRATE) down 

migration_drop: 
		$(MIGRATE) drop -f

migration_status:
		$(MIGRATE) version

migration_force:
		$(MIGRATE) force $(version)

NEWMIGRATE=migrate -path ./db/migrations -database "postgres://postgres:1@localhost:5454/cinemax?sslmode=disable" up 1