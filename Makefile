dev:
	go run cmd/main.go run --watch --config Caddyfile

run:
	./caddy run --watch --config Caddyfile

fmt:
	./caddy fmt -overwrite ./Caddyfile
	#./caddy fmt -overwrite ./Caddyfile.dev
	#./caddy fmt -overwrite ./Caddyfile.live

#build:
#	go build -o caddy ./cmd/main.go

xcaddy:
	# https://caddy.community/t/trying-to-build-caddy-from-source-to-test-2-5-0-beta-with-plugins/15393/2
	# xcaddy build --with github.com/airdb/caddywaf \
	# 	--with github.com/caddy-dns/dnspod

SERVICE := sgw

all: help

help: ## Show help messages
	@echo "Container - ${SERVICE} "
	@echo
	@echo "Usage:\tmake COMMAND"
	@echo
	@echo "Commands:"
	@sed -n '/##/s/\(.*\):.*##/  \1#/p' ${MAKEFILE_LIST} | grep -v "MAKEFILE_LIST" | column -t -c 2 -s '#'

build: ## Build
	docker compose build --no-cache

up: build start

start: ## Create and start containers
	docker compose up -d --force-recreate

stop: ## Stop containers
	docker compose stop

log logs: ## Check logs
	docker compose logs ${SERVICE}

logf: ## Check logs
	docker compose logs -f ${SERVICE}


bash: ## Execute a command in a running container
	docker compose exec ${SERVICE} /bin/sh
