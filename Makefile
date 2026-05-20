#docker compose
.PHONY: up down clean logs
up: 
	@./scripts/docker.sh up -d
down: 
	@./scripts/docker.sh down
clean:
	@./scripts/docker.sh down -v
logs:
	@./scripts/docker.sh logs -f

# terraform s3
.PHONY: s3-init s3-validate s3-plan s3-apply s3-output s3-destroy
s3-init:
	cd ./infrastructures/terraform/s3 && terraform init
s3-validate:
	cd ./infrastructures/terraform/s3 && terraform validate
s3-plan:
	cd ./infrastructures/terraform/s3 && terraform plan
s3-apply:
	cd ./infrastructures/terraform/s3 && terraform apply
s3-output:
	cd ./infrastructures/terraform/s3 && terraform output
s3-destroy:
	cd ./infrastructures/terraform/s3 && terraform destroy

# tools & configuration
.PHONY: install configs
install:
	@./scripts/tools.sh
configs:
	@if [ ! -f configs/config.json ]; then \
		cp configs/config.json.example configs/config.json && echo "Created configs/config.json"; \
	else \
		echo "configs/config.json already exists, skipping..."; \
	fi
	@if [ ! -f configs/flags.json ]; then \
		cp configs/flags.json.example configs/flags.json && echo "Created configs/flags.json"; \
	else \
		echo "configs/flags.json already exists, skipping..."; \
	fi
	@if [ ! -f infrastructures/docker/.env ]; then \
		cp infrastructures/docker/.env.example infrastructures/docker/.env && echo "Created infrastructures/docker/.env"; \
	else \
		echo "infrastructures/docker/.env already exists, skipping..."; \
	fi

# Go
.PHONY: run-api
run-api:
	go run cmd/api/main.go

# Migrate
.PHONY: migrate-create migrate-up migrate-down migrate-drop migrate-force migrate-version
migrate-create:
	@./scripts/migration.sh create $(name)

migrate-up:
	@./scripts/migration.sh up $(args)

migrate-down:
	@./scripts/migration.sh down $(args)

migrate-drop:
	@./scripts/migration.sh drop

migrate-force:
	@./scripts/migration.sh force $(version)

migrate-version:
	@./scripts/migration.sh version


# CONFIG_FILE := config.json

# get_json = $(shell python3 -c "import json; f=open('$(CONFIG_FILE)'); data=json.load(f); print(data$(1))")

# POSTGRES_USER := $(call get_json,['postgres']['user'])
# POSTGRES_PASSWORD := $(call get_json,['postgres']['password'])
# POSTGRES_PORT := $(call get_json,['postgres']['port'])
# POSTGRES_DB := $(call get_json,['postgres']['dbname'])
# REDIS_ADDR := $(call get_json,['redis']['addrs'])

# export POSTGRES_USER POSTGRES_PASSWORD POSTGRES_DB REDIS_ADDR

# COMPOSE_FILE := infrastructures/docker/docker-compose.yml
# DOCKER_COMPOSE := docker compose -f $(COMPOSE_FILE)
# DSN := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

# .PHONY: info
# info:
# 	@echo "PostgreSQL User: $(POSTGRES_USER)"
# 	@echo "PostgreSQL Password: $(POSTGRES_PASSWORD)"
# 	@echo "PostgreSQL Database: $(POSTGRES_DB)"
# 	@echo "Redis Addr: $(REDIS_ADDR)"