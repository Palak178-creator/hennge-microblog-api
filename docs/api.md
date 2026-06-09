# API Documentation

## Base URL

```http
http://localhost:8080
```

---

# Authentication Endpoints

## Register User

### Request

```http
POST /register
```

### Body

```json
{
  "username": "palak",
  "email": "palak@example.com",
  "password": "password123"
}
```

### Success Response

```json
{
  "message": "User registered successfully",
  "token": "<jwt_token>"
}
```

---

## Login User

### Request

```http
POST /login
```

### Body

```json
{
  "email": "palak@example.com",
  "password": "password123"
}
```

### Success Response

```json
{
  "message": "Login successful",
  "token": "<jwt_token>"
}
```

---

# User Endpoints

## Get Profile

### Request

```http
GET /profile
```

### Headers

```http
Authorization: Bearer <jwt_token>
```

### Success Response

```json
{
  "id": 1,
  "username": "palak",
  "email": "palak@example.com"
}
```

---

# Post Endpoints

## Create Post

### Request

```http
POST /posts
```

### Headers

```http
Authorization: Bearer <jwt_token>
```

### Body

```json
{
  "title": "My First Post",
  "content": "Hello World!"
}
```

### Success Response

```json
{
  "message": "Post created successfully"
}
```

---

## Get Posts

### Request

```http
GET /posts
```

### Success Response

```json
[
  {
    "id": 1,
    "title": "My First Post",
    "content": "Hello World!"
  }
]
```

---

# Authentication Flow

1. Register a new account.
2. Login using credentials.
3. Receive JWT token.
4. Include token in Authorization header.
5. Access protected routes.

Example:

```http
Authorization: Bearer <jwt_token>
```
