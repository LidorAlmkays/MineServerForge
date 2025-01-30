# Variables
PROTOC_GEN_GO = $(shell which protoc-gen-go || echo "$(HOME)/go/bin/protoc-gen-go")
PROTOC_GEN_GO_GRPC = $(shell which protoc-gen-go-grpc || echo "$(HOME)/go/bin/protoc-gen-go-grpc")
SQLC_PACKAGE = github.com/sqlc-dev/sqlc/cmd/sqlc

.PHONY: install-tools

install-tools:
	@if [ ! -x "$(PROTOC_GEN_GO)" ]; then \
		echo "Installing protoc-gen-go..."; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	fi
	@if [ ! -x "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "Installing protoc-gen-go-grpc..."; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	fi
	@if ! go list -m $(SQLC_PACKAGE) >/dev/null 2>&1; then \
		echo "Installing sqlc..."; \
		go install $(SQLC_PACKAGE)@latest; \
	fi
	@echo "Tools installed successfully."
