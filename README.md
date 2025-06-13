# Go Hexagonal Architecture Project

A Go project implementing the Hexagonal Architecture (also known as Ports and Adapters) pattern, providing a clean and maintainable structure for building scalable applications.

## Project Structure

```
.
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── adapter/          # Adapters (ports implementation)
│   │   ├── http/        # HTTP API adapters
│   │   ├── entities/    # Adapter-specific entities
│   │   ├── storage/     # Database adapters
│   │   └── config/      # Configuration adapters
│   ├── ports/           # Port interfaces
│   └── core/            # Core business logic
│       ├── entities/    # Domain entities
│       └── service/     # Business services
├── go.mod               # Go module file
├── go.sum               # Go module checksum
├── Makefile            # Build automation
└── .env                # Environment variables (not tracked in git)
```

## Prerequisites

- Go 1.23.1 or higher
- PostgreSQL
- Make (for build automation)

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations
- [squirrel](https://github.com/Masterminds/squirrel) - SQL query builder
- [env](https://github.com/caarlos0/env) - Environment variable parsing
- [godotenv](https://github.com/joho/godotenv) - .env file support

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/thapakazi/go-hex-arch.git
   cd go-hex-arch
   ```

2. Create a `.env` file with the following variables:
   ```
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=5432
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run database migrations:
   ```bash
   make migrateup
   ```

5. Start the application:
   ```bash
   make run
   ```

## Available Make Commands

- `make run` - Run the application
- `make migration name=<migration_name>` - Create a new migration
- `make migrateup` - Run all pending migrations
- `make migrateup1` - Run the next pending migration
- `make migratedown` - Rollback all migrations
- `make migratedown1` - Rollback the last migration
- `make force version=<version>` - Force a specific migration version
- `make docker-migrateup` - Run migrations using Docker

## Architecture

This project follows the Hexagonal Architecture pattern (also known as Ports and Adapters), which is a software design pattern that aims to create loosely coupled application components that can be easily connected to their software environment by means of ports and adapters.

### Core Concepts

1. **Domain Layer (Center)**
   - Contains the core business logic and entities
   - Independent of external concerns
   - Defines the business rules and models
   - Pure business logic with no dependencies on external frameworks

2. **Ports (Interfaces)**
   - **Primary/Driving Ports**: Define how the application is used (e.g., HTTP API endpoints)
   - **Secondary/Driven Ports**: Define how the application interacts with external services (e.g., database operations)
   - Act as contracts between the domain and the outside world
   - Allow the domain to remain independent of external concerns

3. **Adapters (Implementations)**
   - **Primary/Driving Adapters**: Implement the primary ports (e.g., HTTP handlers, CLI commands)
   - **Secondary/Driven Adapters**: Implement the secondary ports (e.g., database repositories, external service clients)
   - Handle the technical details of interacting with the outside world
   - Can be easily swapped without affecting the domain logic

### Benefits

- **Independence**: The domain logic is independent of external concerns
- **Testability**: Easy to test business logic in isolation
- **Flexibility**: External dependencies can be easily swapped
- **Maintainability**: Clear separation of concerns
- **Scalability**: Components can be scaled independently

### Implementation in This Project

```
internal/
├── core/                # Core business logic(Domain)
│   ├── entities/       # Domain entities
│   └── service/        # Business logic services
├── ports/              # Port interfaces
├── adapter/            # Adapters implementation
│   ├── http/          # HTTP API adapters (primary)
│   ├── entities/      # Adapter-specific entities
│   ├── storage/       # Database adapters (secondary)
│   └── config/        # Configuration adapters
```

### Flow of Control

1. External request comes through a primary adapter (e.g., HTTP handler)
2. Primary adapter converts the request to domain objects
3. Domain logic processes the request
4. If needed, domain logic uses secondary ports to interact with external services
5. Secondary adapters implement the actual external interactions
6. Response flows back through the same path



This architecture ensures that:
- Business logic remains pure and independent
- External dependencies can be easily mocked for testing
- The application can be adapted to different external systems
- Changes in external systems don't affect the core business logic

## License

This project is licensed under the MIT License - see the LICENSE file for details. 