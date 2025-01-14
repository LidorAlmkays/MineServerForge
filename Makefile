# Variables
PROTOC_GEN_GO = $(shell which protoc-gen-go || echo "$(HOME)/go/bin/protoc-gen-go")
PROTOC_GEN_GO_GRPC = $(shell which protoc-gen-go-grpc || echo "$(HOME)/go/bin/protoc-gen-go-grpc")
PROTO_DIR = ./internal/api/GRPC/protobufs
PB_DIR = ./internal/api/GRPC/pb
DIST_DIR = ./dist
SRC_DIR = ./cmd

# Targets
.PHONY: all install-tools build clean serve serve-dev serve-prod generate-protos

all: build

install-tools:
	@if [ ! -x "$(PROTOC_GEN_GO)" ]; then \
		echo "Installing protoc-gen-go..."; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	fi
	@if [ ! -x "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "Installing protoc-gen-go-grpc..."; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	fi
	@echo "Tools installed successfully."

generate-protos:
	@echo "Cleaning Protobuf output folder..."
	@rm -rf $(PB_DIR)
	@mkdir -p $(PB_DIR)
	@echo "Generating Protobuf files..."
	@for proto in $(PROTO_DIR)/*.proto; do \
		protoc \
			--proto_path=$(PROTO_DIR) \
			--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
			--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
			--go_out=$(PB_DIR) \
			--go-grpc_out=$(PB_DIR) \
			$$proto; \
	done
	@echo "Protobuf files generated successfully."

build: install-tools generate-protos
	@echo "Building the project..."
	@rm -rf $(DIST_DIR)
	@mkdir -p $(DIST_DIR)
	@go build -o $(DIST_DIR)/app $(SRC_DIR)/main.go
	@echo "Build complete. Executable is in the 'dist' folder."

clean:
	@echo "Cleaning up dist and Protobuf folders..."
	@rm -rf $(DIST_DIR) $(PB_DIR)
	@echo "Cleanup complete."

serve: serve-dev

serve-dev:
	@echo "Starting development server..."
	@go run $(SRC_DIR)/main.go -Mode development

serve-prod: build
	@echo "Starting production server..."
	@$(DIST_DIR)/app -Mode production
