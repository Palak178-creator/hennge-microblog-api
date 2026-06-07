# 🚀 HENNGE Microblog API

A production-style backend REST API built using Go (Gin), PostgreSQL, GORM, and JWT authentication.

---

## 🔥 Features

- User Registration & Login
- JWT Authentication System
- Protected Routes (Middleware)
- Create & Fetch Posts
- User-linked posts (Relational DB)
- PostgreSQL integration with GORM
- Auto database migration

---

## 🛠 Tech Stack

- Go (Gin Framework)
- PostgreSQL
- GORM ORM
- JWT Authentication
- REST API Architecture

---

## 📡 API Endpoints

### Auth
- POST /register
- POST /login

### User
- GET /profile (Protected)

### Posts
- POST /posts (Protected)
- GET /posts

---

## 🔐 Authentication

All protected routes require:
