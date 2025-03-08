# AutoVerse: A Modular Go Framework for RESTful APIs

AutoVerse is a **modular Go framework** designed to simplify building **RESTful APIs** and backend applications. Whether you're building a **microservice** or a **monolithic application**, AutoVerse provides the tools you need to get started quickly and efficiently.

---

## Why Use AutoVerse?

If you're looking for a **scalable**, **modular**, and **production-ready** framework to build RESTful APIs in Go, AutoVerse is for you. Here's why:

- 🛠️ **Built-in Features**: Database migrations, user management, and authentication are included out of the box.
- 🚀 **Quick Start**: Get your API up and running in minutes with a clean and intuitive structure.
- 🧩 **Modular Design**: Easily extend or replace components to fit your specific needs.
- 🗄️ **MySQL Integration**: Seamlessly connect to MySQL for data storage and management.

---

## How Can You Use AutoVerse?

AutoVerse is designed to be flexible and adaptable to your needs. Here are some ways you can use it:

### 1. **As a Template for New Projects**
   - Use AutoVerse as a starting point for your next Go project. Clone the repository, customize the modules, and start building your application.

### 2. **As a Dependency in Existing Projects**
   - Import specific modules (e.g., database migrations or authentication) into your existing Go projects.

### 3. **As a Learning Resource**
   - Explore the codebase to learn best practices for building modular and scalable Go applications.

---

## Features

- **Database Migrations**: Manage schema changes with ease using built-in migration tools.
- **User Management**: Handle user authentication and authorization effortlessly.
- **MySQL Integration**: Connect to MySQL databases seamlessly.
- **Modular Architecture**: Add, remove, or replace modules to suit your application's needs.

---

## Getting Started

### Prerequisites
- Go (version 1.20 or higher)
- MySQL

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Muga20/Go-Modular-Application.git
   cd Go-Modular-Application
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up the database:
   - Update the database configuration in `config/config.go`.
   - Run migrations:
     ```bash
     make migrate-up
     ```

4. Start the application:
   ```bash
   make run
   ```

---

## Running the Application

### Start the Application
To start the application, run:
```bash
make run
```

This will start the application on the default port (e.g., `:8080`).

### Build the Application
To build the application and generate an executable, run:
```bash
make build
```

The executable will be placed in the `bin/` folder.

---

## Database Migrations

### Create a Migration
To create a new migration for a module, use:
```bash
make create-migration MODULE_NAME=users MIGRATION_DESC=create_users_table
```

This generates two migration files in the `Modules/users/migrations/` folder:
- `20231010120000_create_users_table.up.sql`
- `20231010120000_create_users_table.down.sql`

Edit these files to include the necessary SQL for creating and dropping tables.

### Apply Migrations
To apply all migrations, run:
```bash
make migrate-up
```

To apply migrations for a specific module, run:
```bash
make migrate-up-module MODULE_NAME=users
```

### Rollback Migrations
To rollback all migrations, run:
```bash
make migrate-down
```

To rollback migrations for a specific module, run:
```bash
make migrate-down-module MODULE_NAME=users
```

---

## Folder Structure

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

### 2. Database Connection Issues
Ensure the database configuration is correct in `config/config.go`. Verify that the database is running and accessible.

---

## Contributing

We welcome contributions! If you'd like to contribute to AutoVerse, feel free to:
- Open an issue to report bugs or suggest features.
- Submit a pull request with your improvements.

Check out our [Contributing Guidelines](CONTRIBUTING.md) for more details.

---

## License

AutoVerse is licensed under the MIT License. See [LICENSE](LICENSE) for more details.

---

## Questions or Feedback?

If you have any questions or feedback, feel free to open an issue or reach out. Let’s build something awesome together! 🚀

