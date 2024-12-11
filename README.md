# API Project in Go

This repository contains a simple API developed in Go, designed as a learning exercise. The project demonstrates fundamental concepts of backend development, including authentication, CRUD operations, and RESTful design.

## Features

- **JWT Authentication**: Secure token-based authentication for user sessions.
- **CRUD Operations**: Create, read, update, and delete operations for managing products.
- **User Management**: Endpoint for user creation and JWT token generation.
- **Framework**: Built using the [Chi](https://github.com/go-chi/chi) router, a lightweight and idiomatic Go HTTP framework.
- **Swagger Documentation**: Integrated API documentation with Swagger.

## Project Structure

The project is structured to keep the codebase clean and modular:

```
.
├── cmd/
│   └── server/           // Application entry point
│       ├── .env          // Environment variables
│       ├── main.go       // Main application logic
│       └── test.db       // Test database file
├── configs/              // Configuration files
│   └── config.go         // Application configurations
├── docs/                 // API documentation
│   ├── docs.go           // Swagger setup
│   ├── swagger.json      // Swagger JSON spec
│   └── swagger.yaml      // Swagger YAML spec
├── internal/             // Core application logic
│   ├── dto/              // Data Transfer Objects
│   │   └── dto.go
│   └── entity/           // Business entities
│       ├── product.go    // Product entity
│       ├── product_test.go
│       ├── user.go       // User entity
│       └── user_test.go
├── infra/                // Infrastructure-related code
│   ├── database/         // Database interactions
│   └── webserver/        // HTTP server setup
├── pkg/                  // Shared packages
│   └── entity/           // Entity-related utilities
│       └── id.go
├── test/                 // HTTP test files
│   ├── product.http      // Product-related tests
│   └── user.http         // User-related tests
├── go.mod                // Go module definition
├── go.sum                // Dependencies checksum
└── .gitignore            // Git ignore rules
```

## Prerequisites

- **Go**: Version 1.19 or later.
- **Database**: (Optional) This project does not include database integration but uses a test database for development purposes.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Authentication
- **POST /user/generate_token**: Generate a JWT token for an existing user.

### User Management
- **POST /user**: Create a new user.

### Product Management
- **POST /product/**: Add a new product.
- **GET /product/{id}**: Retrieve a specific product by ID.
- **PUT /product/{id}**: Update an existing product by ID.
- **DELETE /product/{id}**: Delete a product by ID.
- **GET /product/**: Retrieve a list of products.

## JWT Authentication

This API uses JWT (JSON Web Tokens) for secure authentication. The token must be included in the `Authorization` header of requests to protected endpoints:

```
Authorization: Bearer <token>
```

## Contribution

This project was created for personal learning and practice, but contributions are welcome. Feel free to fork the repository and submit pull requests.

## License

This project is open-source and available under the MIT License. See the [LICENSE](LICENSE) file for more information.
