app_container := dev_server

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'make dev: make dev for development work'
	@echo 'clean: clean for all clear docker images'
	@echo 'generate-mockery: generate mockery for testing'

dev:
	docker-compose -f docker-compose.yml down
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker-compose -f docker-compose.yml up --build --attach server --attach postgres_db

clean:
	docker-compose -f docker-compose.yml down -v

generate-mockery:
	docker exec -it $(app_container) /bin/sh -c "mockery --all"