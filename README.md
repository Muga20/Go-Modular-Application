# AutoVerse Application Documentation

This document provides a comprehensive guide to setting up, running, and managing the **AutoVerse** application. The application is built using Go and follows a modular approach with support for database migrations.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Running the Application](#running-the-application)
4. [Database Migrations](#database-migrations)
   - [Creating Migrations](#creating-migrations)
   - [Applying Migrations](#applying-migrations)
   - [Rolling Back Migrations](#rolling-back-migrations)
   - [Forcing a Migration Version](#forcing-a-migration-version)
   - [Running Migrations for a Specific Module](#running-migrations-for-a-specific-module)
5. [Folder Structure](#folder-structure)
6. [Configuration](#configuration)
7. [Troubleshooting](#troubleshooting)
8. [Contributing](#contributing)

---

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.20 or higher)
- **MySQL** (or another supported database)
- **Make** (for running commands via `Makefile`)
- **Git** (for version control)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/auto_verse.git
   cd auto_verse
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up the database:
   - Create a MySQL database (e.g., `auto_verse`).
   - Update the database configuration in `config/config.go`.

---

## Running the Application

### 1. Start the Application

To start the application, run:
```bash
make run
```

This will start the application on the default port (e.g., `:8080`).

---

### 2. Build the Application

To build the application and generate an executable, run:
```bash
make build
```

The executable will be placed in the `bin/` folder.

---

## Database Migrations

### Creating Migrations

To create a new migration for a module, use the `create-migration` target:
```bash
make create-migration MODULE_NAME=users MIGRATION_DESC=create_users_table
```

This will generate two migration files in the `Modules/users/migrations/` folder:
- `20231010120000_create_users_table.up.sql`
- `20231010120000_create_users_table.down.sql`

Edit these files to include the necessary SQL for creating and dropping tables.

---

### Applying Migrations

To apply all migrations, run:
```bash
make migrate-up
```

This will execute the SQL in the `up` migration files for all modules.

To apply migrations for a specific module, run:
```bash
make migrate-up-module MODULE_NAME=users
```

This will execute the SQL in the `up` migration files for the specified module.

---

### Rolling Back Migrations

To rollback all migrations, run:
```bash
make migrate-down
```

This will execute the SQL in the `down` migration files for all modules.

To rollback migrations for a specific module, run:
```bash
make migrate-down-module MODULE_NAME=users
```

This will execute the SQL in the `down` migration files for the specified module.

---

### Running Migrations for a Specific Module

You can apply or rollback migrations for a specific module using the following commands:

1. **Apply Migrations for a Specific Module:**
   ```bash
   make migrate-up-module MODULE_NAME=users
   ```

2. **Rollback Migrations for a Specific Module:**
   ```bash
   make migrate-down-module MODULE_NAME=users
   ```

---

## Folder Structure

The project follows this structure:

```
auto_verse/
├── bin/                     # Compiled executable
├── cmd/                     # Application entry points
│   └── main.go
├── config/                  # Configuration files
│   └── config.go
├── helpers/                 # Helper scripts
│   └── create_module.go
├── migrations/              # Migration management
│   └── registry.go
├── Modules/                 # Application modules
│   └── users/               # Example module
│       ├── migrations/      # Migration files for the module
│       ├── models/          # Database models
│       ├── routes/          # HTTP routes
│       └── ...
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksum
├── Makefile                 # Makefile for running commands
└── README.md                # Project documentation
```

---

## Configuration

The application configuration is stored in `config/config.go`. Update the following environment variables:

```go
var Envs = struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}{
	DBUser:     os.Getenv("DB_USER"),
	DBPassword: os.Getenv("DB_PASSWORD"),
	DBAddress:  os.Getenv("DB_ADDRESS"),
	DBName:     os.Getenv("DB_NAME"),
}
```

Set these variables in your environment or in a `.env` file.

---

## Troubleshooting

### 1. Dirty Database Version

If you encounter the error `Dirty database version`, follow these steps:

1. Identify the failed migration:
   ```bash
   cat Modules/users/migrations/20231010120000_create_users_table.up.sql
   ```

2. Fix the SQL in the migration file.

3. Force the database version:
   ```bash
   make force-version VERSION=20231010120000
   ```

4. Apply migrations again:
   ```bash
   make migrate-up
   ```

---

### 2. Database Connection Issues

Ensure the database configuration is correct in `config/config.go`. Verify that the database is running and accessible.

---

## Contributing

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---


Thank you for using **AutoVerse**! 🚀

---
