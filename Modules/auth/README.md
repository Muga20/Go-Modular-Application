# Auth Module

This module handles auth-related functionality.

## Structure
- **controllers/**: Handles HTTP requests.
- **models/**: Defines database models.
- **routes/**: Defines API routes.
- **middleware/**: Middleware for the module.
- **utils/**: Utility functions.
- **config/**: Module-specific configuration.
- **tests/**: Unit and integration tests.
- **migrations/**: Database migration files.

## Usage
- To use the module, import it in your application and call the setup function in your main file.

## Migrations
- Run migrations using the `Migrate` function in `migrate.go`.

## Testing
- Run `go test ./Modules/auth/tests` to run tests for this module.
