# JWT-Authenticated Microblog API

A secure RESTful microblogging backend built with Go, featuring JWT-based authentication, protected routes, Docker containerization, and a modular project structure.

## Features

- User Registration
- User Login with JWT Authentication
- Protected API Routes
- Create and View Posts
- Authorization Middleware
- Docker Support
- Modular Project Architecture

## Tech Stack

- Go
- JWT (JSON Web Tokens)
- Docker
- Docker Compose
- REST API

## Project Structure

```text
hennge-microblog-api
│
├── cmd
│   └── api
│       └── main.go
│
├── config
│   └── db.go
│
├── internal
│   ├── handlers
│   ├── middleware
│   └── models
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

### Authentication

| Method | Endpoint  | Description                 |
| ------ | --------- | --------------------------- |
| POST   | /register | Register a new user         |
| POST   | /login    | Login and receive JWT token |

### Posts

| Method | Endpoint | Description                       |
| ------ | -------- | --------------------------------- |
| GET    | /posts   | Get all posts                     |
| POST   | /posts   | Create a new post (Authenticated) |

### Users

| Method | Endpoint    | Description          |
| ------ | ----------- | -------------------- |
| GET    | /users/{id} | Get user information |

## Authentication Flow

1. User registers an account.
2. User logs in with credentials.
3. Server generates a JWT token.
4. Client sends the JWT token in the Authorization header.
5. Middleware validates the token before allowing access to protected routes.

## Running Locally

### Clone Repository

```bash
git clone <repository-url>
cd hennge-microblog-api
```

### Run with Go

```bash
go run cmd/api/main.go
```

### Run with Docker

```bash
docker-compose up --build
```

## Future Improvements

- PostgreSQL Integration
- Refresh Tokens
- Rate Limiting
- API Documentation (Swagger)
- Unit Testing
- CI/CD Pipeline

## Author

Palak Patel

BTech ICT, DAIICT
