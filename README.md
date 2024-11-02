# Netly

Netly is a simple social network API.

## Features

- User registration and authentication
- Create, read, update, and delete (CRUD) user profiles
- Post and comment functionality
- API documentation with Swagger
- Seamless database migrations with Golang Migrate

## Tech Stack

- **Go** - Primary programming language
- **Docker** - Containerization for easy setup and deployment
- **PostgreSQL** - Relational database management system
- **Swagger** - API documentation
- **Golang Migrate** - Database migrations

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/bjclayton/netly.git
cd netly
```

### 2. Configure Environment Variables

Create an `.envrc` file in the root directory and add the following environment variables:

```plaintext
DB_DRIVER=postgres
DB_ADDR=postgres://root:secret@localhost:5432/gosocialnet?sslmode=disable
SERVER_ADDRESS=:5000
MIGRATIONS_PATH=./cmd/migrate/migrations
```

Load the `.envrc` file by running:

```bash
direnv allow
```

### 3. Start the Application with Docker

To run the app and database in Docker containers, use the following command:

```bash
docker-compose up
```

This will start the application on `http://localhost:5000`.

### 4. Run Database Migrations

Use the Makefile commands to manage database migrations.

- **Create a new migration**:
  ```bash
  make migration name=<migration_name>
  ```

- **Apply migrations**:
  ```bash
  make migrate-up
  ```

- **Rollback migrations**:
  ```bash
  make migrate-down
  ```

### 5. Seed the Database

To run a database seeding script:

```bash
make seed
```

### 6. Run Tests

To run all tests with verbose output:

```bash
make test
```

### 7. Generate API Documentation

To generate Swagger documentation:

```bash
make gen-docs
```