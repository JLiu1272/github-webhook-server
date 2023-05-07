say_hello:
	@echo "Hello World"

build: 
	@echo "Building the project"
	@go build -o bin/$(BINARY_NAME) -v

run: 
	@go run server.go