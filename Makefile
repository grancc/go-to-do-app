include .env
export

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Clean all volumes? [y/N]" ans; \ 
	if [ "$$ans" = "y" ]; then \
		docke composew down todoapp-postgres && \
		rm -rf out/pg_data && \
		echo "Done!"; \
	else \
		echo "Cleanup is cancel"; \
	fi

migrate-create:
	@if [-z "$(seq)" ]; then \ 
		echo "Missing requirement parametr seq. Example: migrate-create seq=init"; \
		exit 1;
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [-z "$(action)" ]; then \ 
		echo "Missing requirement parametr action. Example: migrate-create action=up"; \
		exit 1;
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

migrate-up:
	@migrate-action action=up

migrate-down:
	@migrate-action action=down