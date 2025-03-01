# URL Shortening Service

This project is a URL shortening service, inspired by the backend projects section of [roadmap.sh](https://roadmap.sh/projects/url-shortening-service).
The project is built using Clean Architecture and various design patterns to ensure maintainability and scalability.
The API documentation is provided using Swagger, and the entire project is containerized using Docker.

## Features
- Shorten long URLs
- Redirect to original URLs using short codes
- Track access counts for each short URL
- API documentation with Swagger
- Clean Architecture for maintainability
- Dockerized for easy deployment

## Architecture
The project follows the principles of Clean Architecture, ensuring that the core business logic is decoupled from the infrastructure and frameworks. The main layers are:
- **Entities**: Core business objects
- **Use Cases**: Application-specific business rules
- **Interface Adapters**: Converters and presenters
- **Frameworks & Drivers**: External interfaces like databases, web frameworks, etc.

## Design Patterns
- **Repository Pattern**: For data access abstraction
- **Dependency Injection**: For managing dependencies
- **DTO (Data Transfer Object)**: For transferring data between layers

## API Documentation
The API is documented using Swagger. You can access the Swagger UI at `/swagger/index.html` once the application is running.

## Docker Setup
The project is containerized using Docker. The `docker-compose.yml` file sets up the following services:
- **Backend**: The main application
- **Postgres**: The database
- **PgAdmin**: Database management tool
- **Redis**: In-memory data structure store for caching

## Getting Started
1. **Clone the repository**:
    ```sh
    git clone https://github.com/alielmi98/go-url-shortener.git
    cd go-url-shortener
    ```

2. **Build and run the Docker containers**:
    ```sh
    cd docker
    docker-compose up --build
    ```

3. **Access the services**:
    - **Backend**: `http://localhost:5000`
    - **PgAdmin**: `http://localhost:8090`
    - **Swagger UI**: `http://localhost:5000/swagger/index.html`

## Contributing
Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvements.

## License
This project is licensed under the MIT License.