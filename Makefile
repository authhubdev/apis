.PHONY: build
build:
	@echo "Building..."
	@buf lint && buf generate
	@echo "Build complete"
	@go mod tidy
