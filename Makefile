.PHONY: buid
build:
	@echo "Building..."
	@buf lint && buf generate
	@echo "Build complete"
