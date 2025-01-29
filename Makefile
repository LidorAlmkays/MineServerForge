include Makefile.tools.mk
include internal/api/Makefile.protobufs.mk
include db/Makefile.sqlc.mk

# Variables
SWAGGER_OUT_DIR = ./internal/api/docs
DIST_DIR = ./dist
SRC_DIR = ./cmd

.PHONY: all serve serve-dev serve-prod build clean

all: build

serve: serve-dev

serve-dev: build
	@echo "Initializing Swagger documentation..."
	@swag init -g $(SRC_DIR)/main.go -o $(SWAGGER_OUT_DIR)
	@echo "Starting development server..."
	@go run $(SRC_DIR)/main.go -Mode development

serve-prod: build
	@echo "Starting production server..."
	@$(DIST_DIR)/app -Mode production

build: install-tools generate-protos generate-sqlc
	@echo "Building the project..."
	@rm -rf $(DIST_DIR)
	@mkdir -p $(DIST_DIR)
	@go build -o $(DIST_DIR)/app $(SRC_DIR)/main.go
	@echo "Build complete. Executable is in the 'dist' folder."

clean:
	@echo "Cleaning up dist, SQLC, and Protobuf folders..."
	@rm -rf $(DIST_DIR) $(PB_DIR)
	@echo "Cleanup complete."
