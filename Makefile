run: build
	@./bin/hyperchat

build: 
	@go build -o bin/hyperchat cmd/server/main.go

test: 
	@go test -v ./...
