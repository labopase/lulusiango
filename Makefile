#docker compose
.PHONY: up down clean logs
up:
	@./scripts/docker-compose.sh up -d
down:
	@./scripts/docker-compose.sh down
clean:
	@./scripts/docker-compose.sh down -v
logs:
	@./scripts/docker-compose.sh logs -f

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
		cp infrastructures/docker-compose/.env.example infrastructures/docker-compose/.env && echo "Created infrastructures/docker/.env"; \
	else \
		echo "infrastructures/docker-compose/.env already exists, skipping..."; \
	fi

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

# SQLC
.PHONY: sqlc-gen sqlc-clean copy-sqlc
sqlc-gen:
	@echo "Generating SQLC code..."
	@sqlc generate

sqlc-clean:
	@echo "Cleaning generated SQLC code..."
	@rm -f database/sqlc/models.go database/sqlc/db.go database/sqlc/querier.go database/sqlc/*.sql.go

copy-sqlc:
	@cp /Users/mymac/.gvm/pkgsets/go1.25.3/global/bin/sqlc ./sqlc
	@chmod +x ./sqlc