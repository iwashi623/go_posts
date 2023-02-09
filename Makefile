.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build the docker image
	docker build -t budougumi0617/gotodo:$(DOCKER_TAG) \
	--target deploy ./

build-local: ## Build the docker image for local
	docker-compose build --no-cache

up: ## Run the docker container
	docker-compose up -d

down: ## Stop the docker container
	docker-compose down

logs: ## Show the docker container logs
	docker-compose logs -f

ps: ## Show the docker container status
	docker-compose ps

test: ## Run the test
	go test -race -shuffle=on ./...

help: ## Show option
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
