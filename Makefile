# Makefile for auto_verse project

# Variables
APP_NAME=auto_verse
MODULES_DIR=Modules
BIN_DIR=bin

# Default target
all: run

# Run the application
run:
	@echo "Starting the application..."
	go run cmd/main.go

# Build the application and place the executable in ./bin/
build:
	@echo "Building the application..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/main.go
	@echo "Build complete. Executable: $(BIN_DIR)/$(APP_NAME)"

# Generate a new module
create-module:
ifndef MODULE_NAME
	$(error MODULE_NAME is not set. Usage: make create-module MODULE_NAME=<module_name>)
endif
	@echo "Creating new module: $(MODULE_NAME)"
	go run helpers/create_module.go $(MODULE_NAME)
	@echo "Module '$(MODULE_NAME)' created in $(MODULES_DIR)/$(MODULE_NAME)"

# Generate a new migration
create-migration:
ifndef MODULE_NAME
	$(error MODULE_NAME is not set. Usage: make create-migration MODULE_NAME=<module_name> MIGRATION_DESC=<description>)
endif
ifndef MIGRATION_DESC
	$(error MIGRATION_DESC is not set. Usage: make create-migration MODULE_NAME=<module_name> MIGRATION_DESC=<description>)
endif
	@echo "Creating new migration for module: $(MODULE_NAME)"
	go run cmd/create_migration/main.go $(MODULE_NAME) $(MIGRATION_DESC)
	@echo "Migration created in $(MODULES_DIR)/$(MODULE_NAME)/migrations/"

# Run migrations (up) for all modules
migrate-up:
	@echo "Applying migrations (up)..."
	go run cmd/main.go --migrate up
	@echo "Migrations applied successfully!"

# Run migrations (up) for a specific module
migrate-up-module:
ifndef MODULE_NAME
	$(error MODULE_NAME is not set. Usage: make migrate-up-module MODULE_NAME=<module_name>)
endif
	@echo "Applying migrations (up) for module: $(MODULE_NAME)..."
	go run cmd/main.go --migrate up --module $(MODULE_NAME)
	@echo "Migrations applied successfully for module: $(MODULE_NAME)!"

# Rollback migrations (down) for all modules
migrate-down:
	@echo "Rolling back migrations (down)..."
	go run cmd/main.go --migrate down
	@echo "Migrations rolled back successfully!"

# Rollback migrations (down) for a specific module
migrate-down-module:
ifndef MODULE_NAME
	$(error MODULE_NAME is not set. Usage: make migrate-down-module MODULE_NAME=<module_name>)
endif
	@echo "Rolling back migrations (down) for module: $(MODULE_NAME)..."
	go run cmd/main.go --migrate down --module $(MODULE_NAME)
	@echo "Migrations rolled back successfully for module: $(MODULE_NAME)!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)
	@echo "Clean complete."

# Help: Display available targets
help:
	@echo "Available targets:"
	@echo "  run               - Start the application"
	@echo "  build             - Build the application and place the executable in ./bin/"
	@echo "  create-module     - Generate a new module (Usage: make create-module MODULE_NAME=<module_name>)"
	@echo "  create-migration  - Generate a new migration (Usage: make create-migration MODULE_NAME=<module_name> MIGRATION_DESC=<description>)"
	@echo "  migrate-up        - Apply database migrations (up) for all modules"
	@echo "  migrate-up-module - Apply database migrations (up) for a specific module (Usage: make migrate-up-module MODULE_NAME=<module_name>)"
	@echo "  migrate-down      - Rollback database migrations (down) for all modules"
	@echo "  migrate-down-module - Rollback database migrations (down) for a specific module (Usage: make migrate-down-module MODULE_NAME=<module_name>)"
	@echo "  clean             - Remove build artifacts"
	@echo "  help              - Display this help message"