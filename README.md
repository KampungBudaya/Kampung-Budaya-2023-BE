<img src="https://raw.githubusercontent.com/MirzaHilmi/MirzaHilmi/master/art/kampung-budaya-2023.png" alt="Kampung Budaya 2023 Logo">

# 2023 Kampung Budaya Backend Service

This repository contains the backend service for the Kampung Budaya Festival, built using the Golang programming language and the Gorilla Mux router. The service also utilizes Swagger for API documentation, powered by the Swaggo/Swag module. Additionally, the project includes a database migration setup managed through the go-migrate module and controlled using the provided Makefile.

## Getting Started

Follow these instructions to set up and run the backend service locally.

### Prerequisites

- Golang must be installed on your machine. If not, you can download it from the [official Golang website](https://golang.org/dl/).
- Ensure you have a MySQL/MariaDB database available.
- Make package to run make automated commands

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/KampungBudaya/Kampung-Budaya-2023-BE.git
   cd Kampung-Budaya-2023-BE
   ```

2. Copy `.env.example` file to `.env` in the project root directory and set your database credentials:

   ```bash
   cp .env.example .env
   ```

3. Install project dependencies:

   ```bash
   go mod download
   ```

4. Run the database migrations using the provided Makefile:

   ```bash
   make migrate-up
   ```

5. Build and run the service:

   ```bash
   go run main.go
   ```

The service should now be running locally at `http://localhost:8080` or any port you desire to use.

## API Documentation

The API documentation for the backend service can be accessed through Swagger UI. After starting the service locally, navigate to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to view and interact with the API endpoints.

## Database Migrations

The database migration scripts are managed using the go-migrate module. You can use the provided Makefile commands to handle migrations.

### Available Makefile Commands

- `make migrate-up`: Apply database migrations.
- `make migrate-down`: Drop all migration tables.
- `make migrate-drop`: Drop entire schema's table.
- `make migrate-version`: Display the current migration version.

## Contributing

Contributions are welcome! If you find any issues or have improvements to suggest, please feel free to open an issue or submit a pull request.