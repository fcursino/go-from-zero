
# Go API with Gin-Gonic and PostgreSQL

This repository contains an API developed in Go using the Gin-Gonic framework. The API performs CRUD (Create, Read, Update, Delete) operations on a PostgreSQL database and implements authentication using JWT (JSON Web Tokens).

## Features

- **CRUD**: Create, read, update, and delete resources.
- **JWT Authentication**: Protects API routes, ensuring that only authenticated users can access certain functionalities.
- **Database**: Uses PostgreSQL for data storage.

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin-Gonic](https://gin-gonic.com/) - Web framework for Go
- [PostgreSQL](https://www.postgresql.org/) - Database management system
- [JWT](https://jwt.io/) - Authentication standard

## Prerequisites

Before you begin, you will need to have the following installed on your machine:

- Go (version 1.24 or higher)
- PostgreSQL
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/fcursino/go-from-zero.git
   cd go-from-zero
   ```
2. Build and run the application using Docker:

   ```bash
   docker-compose up --build
   ```
3. The API will be available at `http://localhost:8000`.

## Docker Configuration

The `docker-compose.yml` file defines two services: `go_db` for the PostgreSQL database and `go-app` for the Go application.

```yaml
version: "3.9"

services: 
  go_db:
      container_name: go_db
      image: postgres:12
      environment:
        POSTGRES_PASSWORD: 1234
        POSTGRES_USER: postgres
        POSTGRES_DB: postgres
      ports:
        - "5432:5432"
      volumes:
        - pgdata:/var/lib/postgresql/data
  go-app:
    container_name: go-app
    image: go-from-zero-api
    build: .
    ports:
      - "8000:8000"
    depends_on:
      - go_db
volumes:
  pgdata: {}
```
### Explanation of Docker Services
go_db: This service runs a PostgreSQL database. It sets the database user, password, and database name through environment variables. The database is accessible on port 5432.
go-app: This service runs the Go application. It builds the application from the current directory and exposes it on port 8000. It depends on the go_db service, ensuring that the database is started before the application.