include .env
export $(shell sed 's/=.*//' .env)

BACKEND := backend

.PHONY: docs run build migration psql stop

run:
	@cd $(BACKEND) && go run ./cmd/test-service/main.go &
	@docker compose up --attach backend

build: docs
	@cd $(BACKEND) && go run ./cmd/test-service/main.go &
	@docker compose up --build --attach backend

stop:
	@docker stop tsso-back
	@docker stop tsso-front
	@fuser -k 9102/tcp

docs:
	@cd $(BACKEND) && swag init -g ./cmd/sso/main.go -o ./docs

migration:
	@docker exec -it sso-db psql -U ${POSTGRES_USER} -d sso -f $(m)

psql:
	@docker exec -it sso-db psql -U ${POSTGRES_USER} -d sso

test:
	@cd $(BACKEND) && go test -v ./...
