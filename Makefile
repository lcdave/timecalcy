include .env

export $(shell sed 's/=.*//' .env)

$(info environment: $(GO_ENV))

.PHONY: migrate
# set default target
.DEFAULT_GOAL := help

help:
	@echo "Make Routines:"
	@echo " - build                            build web containers"
	@echo " - up                               start the services"
	@echo " - down                             stop the services"
	@echo " - remove                           stop and remove containers"

all:
	@echo "🖥️  Starting db container";
	docker-compose -f docker/docker-compose.yml -p timecalcy up -t 2

down:
	@echo "⚠️  Taking db container down"
	docker-compose -f docker/docker-compose.yml -p timecalcy down -t 2

migrate:
	migrate -source file://backend/migrations \
			-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable up

remove: down
	@echo "⚠️  Removing docker containers"
	docker rm db
	docker rm web