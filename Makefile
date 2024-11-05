CURRENT_DIR := $(shell pwd)

run:
	go run cmd/main.go


SWAGGER := $(HOME)/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/api/router.go -o $(SWAGGER_DOCS)

swag-gen:
	@echo "Running: $(SWAGGER_INIT)"
	$(SWAGGER_INIT)