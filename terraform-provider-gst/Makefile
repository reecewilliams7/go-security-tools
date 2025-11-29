.PHONY: build install test clean fmt lint

# Build the provider
build:
	go build -o terraform-provider-gst

# Install dependencies
install:
	go mod download
	go mod tidy

# Run tests
test:
	go test ./... -v

# Clean build artifacts
clean:
	rm -f terraform-provider-gst
	rm -rf examples/test/.terraform*
	rm -rf examples/test/terraform.tfstate*

# Format code
fmt:
	go fmt ./...
	terraform fmt -recursive ./examples/

# Run linters
lint:
	go vet ./...

# Generate documentation
docs:
	go generate ./...

# Run all checks before commit
check: fmt lint test

# Setup local development override
dev-override:
	@echo "Add this to your ~/.terraformrc:"
	@echo ""
	@echo "provider_installation {"
	@echo "  dev_overrides {"
	@echo "    \"reecewilliams7/gst\" = \"$(PWD)\""
	@echo "  }"
	@echo "  direct {}"
	@echo "}"
