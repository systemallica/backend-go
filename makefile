all: build docs start 
build:
	go build -o bin/server ./cmd/server
test:
	go test ./...
start:
	export $$(cat .env | grep -v ^\# | xargs) && ./bin/server
docs:
	swag init -d "cmd/server/,rides/,api/,api/handlers/"
migrate: 
	rel migrate
format: 
	goimports -w ./..
lint:
	golangci-lint run