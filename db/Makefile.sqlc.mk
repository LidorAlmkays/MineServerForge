.PHONY: generate-sqlc

generate-sqlc:
	@echo "Generating SQLC queries..."
	@sqlc generate -f $(CURDIR)/db/sqlc.yaml
	@echo "SQLC queries generated successfully."