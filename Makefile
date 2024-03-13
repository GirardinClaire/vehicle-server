IMAGE?=girardinclaire/vehicle-server
TAG?=dev
DB_CONTAINER_NAME=vehicle-server-devgo
POSTGRES_USER=vehicle-server
POSTGRES_PASSWORD=secret
POSTGRES_DB=vehicle-server
DATABASE_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)


.PHONY: all
all: clean dist unit_test integration_test build package

.PHONY: clean
clean:
	rm -r -f ./dist

.PHONY: dist
dist:
	mkdir -p ./dist

.PHONY: build
build:
	go build -o ./dist/server ./cmd/server/


.PHONY: unit_test
unit_test:
	go test -v -cover ./...

.PHONY: integration_test
integration_test:
	go test -v -count=1 --tags=integration ./app

.PHONY: package
package:
	docker build -t $(IMAGE):$(TAG) .

.PHONY: dev
dev: dev_db
	go run ./cmd/server \
		-listen-address=:8080 \
		-database-url=$(DATABASE_URL)

.PHONY: dev_db
dev_db:
	docker container run \
		--detach \
		--rm \
		--name=$(DB_CONTAINER_NAME) \
		--env=POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		--env=POSTGRES_USER=$(POSTGRES_USER) \
		--env=POSTGRES_DB=$(POSTGRES_DB) \
		--publish 5432:5432 \
		postgis/postgis:16-3.4-alpine

.PHONY: stop_dev_db
stop_dev_db:
	docker container stop $(DB_CONTAINER_NAME)
