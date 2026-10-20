.PHONY: api web dev build run-bin

# Run the Go API
api: 
	cd api && go run ./cmd/api

# Build the Go API
build:
	cd api && go build -o bin/api ./cmd/api

# Execute the Go API binary
run-bin:
	./api/bin/api

# Build and exceute binary with one command
build-run: build run-bin

# Run the React + Vite frontend
web:
	cd app && pnpm dev

# Run both apps concurrently
dev:
	$(MAKE) -j2 api web