# URL Shortener Service

## Overview

This project is a URL shortener service developed using the Golang Fiber framework, featuring built-in rate limiting and containerized deployment via Docker.

## Features

- **Service Development**: Implemented a high-performance URL shortener service using the Golang Fiber framework, known for its efficiency and scalability.
- **Rate Limiting**: Added rate limiting to control and prevent misuse, ensuring fair and consistent service usage.
- **Containerization**: The entire application is containerized into a Docker image, simplifying deployment and enabling scalability.
- **Security Measures**: Followed best practices to enhance the security and reliability of the service.

## Getting Started

To run the URL shortener service locally, follow these steps:

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine
- [Docker](https://docs.docker.com/get-docker/) installed on your machine

### Running the Application Locally

1. **Clone the Repository**

    ```bash
    git clone https://github.com/yourusername/url-shortener-service.git
    cd url-shortener-service
    ```

2. **Build and Run with Docker**

    ```bash
    docker build -t url-shortener-service .
    docker run -p 8080:8080 url-shortener-service
    ```

3. **Access the Service**

    Open your web browser and navigate to `http://localhost:3000` to use the URL shortener service.

### Running the Application With My Pre-built Docker Image

If you prefer to use the pre-built Docker image available on Docker Hub, follow these steps:

1. **Pull the Docker Image**

    ```bash
    docker pull rohanyh/url-shortner
    ```

2. **Run the Docker Container**

    ```bash
    docker run -p 8080:8080 rohanyh/url-shortner
    ```

3. **Access the Service**

    Open your web browser and navigate to `http://localhost:8080` to use the URL shortener service.


### Running the Application Without Docker

1. **Install Dependencies**

    ```bash
    go mod tidy
    ```

2. **Run the Application**

    ```bash
    go run main.go
    ```

3. **Access the Service**

    Open your web browser and navigate to `http://localhost:3000` to start using the URL shortener service.

## API Endpoints

- `POST /api/v1` - Shorten a given URL
- test command:
```powershell
# request
curl --location --request POST 'localhost:3000/api/v1' \
--header 'Content-Type: application/json'
--data-raw { "URL": "https://github.com/rohanyh101" }

# response
{
 "url": "https://github.com/rohanyh101",
 "short": "localhost:3000/w234e90",
 "expiry": 24,
 "rate_limit": 9,
 "rate_limit_reset": 30,
}
```

- `GET /:url` - Retrieve the original URL from the shortened URL
- test command:
```powershell
# request
curl --location --request GET 'localhost:3000/w234e90' \
--header 'Content-Type: application/json'

# response
will redirect to the original URL, i.e. my GitHub page.
```

## Contributing

Feel free to submit issues or pull requests if you have suggestions or improvements. Please adhere to the project's coding standards and guidelines.

## Acknowledgements

- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [Docker](https://www.docker.com/)

