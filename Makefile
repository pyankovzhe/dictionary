.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

.PHONY: check
check:
	golangci-lint run

.PHONY: dbmigrate
dbmigrate:
	@(goose --dir migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=${db} sslmode=disable" up)

.DEFAULT_GOAL := build
