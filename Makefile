include .env
export $(shell sed 's/=.*//' .env)

.PHONY: docs run build migration psql stop

run:
	@docker compose up --attach sso

build: docs
	@go run ./cmd/test-service/main.go &
	@docker compose up --build --attach sso

stop:
	@fuser -k 9102/tcp
	@docker stop toplivo-sso

docs:
	@swag init -g ./cmd/sso/main.go -o ./docs

migration:
	@docker exec -it sso-db psql -U ${POSTGRES_USER} -d sso -f $(m)

psql:
	@docker exec -it sso-db psql -U ${POSTGRES_USER} -d sso

test:
	@go test ./...
