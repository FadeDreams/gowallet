## Wallet Application (Go, Hexagonal, DDD, Docker, PostgreSQL, SQLx, Gorm, JWT)

The Wallet Application is a Go-based system designed with the Hexagonal pattern and modular architecture following Domain-Driven Design (DDD) principles. The application utilizes Gorilla Mux for RESTful API creation, connecting HTTP handlers as adapters to the core logic encapsulated in the `domain` and `service` packages.

Database access is managed through the adapter layer, connecting to PostgreSQL using SQLx with the option of using Gorm. The application ensures secure user access with JWT-based authentication and authorization. Additionally, test cases are included to validate the functionality and reliability of the application.

## Project Structure
```plaintext
/gowallet
|-- app
|-- docker-compose.yml
|-- Dockerfile
|-- domain
|-- dto
|-- errors
|-- go.mod
|-- go.sum
|-- initdb
|-- main.go
|-- service
|-- tmp
```

## Getting Started
1. Clone the repository.
2. Run the Docker Compose configuration using `docker-compose up`.
3. Explore the various microservices and their functionalities.

Feel free to contribute, report issues, or provide feedback. Let's collaborate to further enhance and optimize the Wallet Application!
