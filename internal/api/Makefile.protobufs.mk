# Variables
PROTO_DIR = ./internal/api/GRPC/protobufs
PB_DIR = ./internal/api/GRPC/pb
PROTOC_GEN_GO = $(shell which protoc-gen-go || echo "$(HOME)/go/bin/protoc-gen-go")
PROTOC_GEN_GO_GRPC = $(shell which protoc-gen-go-grpc || echo "$(HOME)/go/bin/protoc-gen-go-grpc")

.PHONY: generate-protos

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
