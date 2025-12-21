
# Run the application locally
run:
	@echo "🚀 Starting containers and running the Go application..."
	@docker-compose up -d
	@sleep 2
	@go run ./cmd/server/main.go

# Shut down containers and remove persistent volumes
reset:
	@echo "🧹 Resetting environment (containers + volumes)..."
	@docker-compose down -v

# proto-gen: Generate Go code from a .proto file
# Usage:
#   make proto-gen PROTO=exchange/create/exchange.proto
# The generated files will be placed in the pkg directory.
proto-gen:
	@echo "🔧 Generating Go code from proto file: $(PROTO)"
	@protoc --go_out=pkg --go_opt=paths=source_relative \
        --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
        proto/$(PROTO)

proto-gen-all:
	@echo "🔧 Generating Go code from all proto files in ./proto..."
	@find ./proto -name '*.proto' | xargs -I {} protoc --go_out=pkg --go_opt=paths=source_relative \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative {}