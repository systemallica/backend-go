all: build docs start 
build:
	go build -o bin/server ./cmd/server
test:
	go test ./...
start:
	export $$(cat .env | grep -v ^\# | xargs) && ./bin/server
docs:
	swag i -g http.go -d "api/,cmd/server/,rides/,api/handlers/"
migrate: 
	rel migrate
format: 
	goimports -w ./..
lint:
	golangci-lint run