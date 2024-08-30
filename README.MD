# Softlancer API

Softlancer is a web application where users can search for projects, apply for them, and log in or register using a Google account. This API is built using the Gin framework in Go, with a project structure inspired by Laravel.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Contributing](#contributing)
- [License](#license)

## Features

- CRUD Projects for Admin
- CRUD Companies for Admin
- Search for projects.
- Apply for projects.

## Project Structure

The project structure follows a pattern inspired by Laravel to keep the application organized and maintainable. The main components of the structure include:

- **`/config`**: Configuration files for the application.
- **`/database`**: Database connection and migration files.
- **`/models`**: Data models representing the database tables.
- **`/routes`**: API route definitions.
- **`/controllers`**: Business logic for handling requests.
- **`/middleware`**: Middleware for request handling, such as authentication.
- **`/utils`**: Utility functions and helpers.
- **`main.go`**: The entry point of the application.

## Setup and Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/softlancer-api.git

   ```

2. Run Migrations

```
go run cmd/database/migration.go
```

3. Run the Application

```
go run cmd/server/main.go
```

Or

```
go run github.com/pilu/fresh
```
