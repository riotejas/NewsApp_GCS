# Simple Makefile

# Build the application
all: build

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build: vet
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Build Docker container
docker-build-local:
	@docker build -f Dockerfile-local --rm -t newsapp:v1.0 .

docker-build-gcp:
	@docker build -f Dockerfile-gcp --rm -t newsapp:v1.0 .

docker-run:
	@docker run -p 8080:8080 newsapp:v1.0

docker-deploy:
	@docker build --platform linux/amd64 --tag gcr.io/news-434519 .

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


.PHONY: all build run test clean
