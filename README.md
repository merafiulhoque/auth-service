<div align="center">

# 🔐 Auth Service

**A standalone, production-ready authentication microservice written in Go.**

Signup · Signin · Email OTP · JWT Access & Refresh Tokens · Magic-link Password Reset

<br />

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Upstash_Redis-00E9A3?style=for-the-badge&logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)
![Resend](https://img.shields.io/badge/Resend-000000?style=for-the-badge&logo=maildotru&logoColor=white)

</div>

---

## 📖 Table of Contents

- [✨ Features](#-features)
- [🧰 Tech Stack](#-tech-stack)
- [🗺️ API Overview](#️-api-overview)
- [📦 Response Format](#-response-format)
- [🔑 Authentication](#-authentication)
- [📡 Endpoint Reference](#-endpoint-reference)
- [🔄 Flows](#-flows)
- [🚀 Getting Started](#-getting-started)
- [⚙️ Environment Variables](#️-environment-variables)
- [🐳 Docker](#-docker)
- [🤝 Contributing](#-contributing)

---

## ✨ Features

- 🧑‍💻 **Email + password auth** with secure **bcrypt** hashing
- 🎟️ **JWT access & refresh tokens** for stateless sessions
- 📧 **Email verification** via 6-digit OTP delivered by **Resend**
- 🪄 **Password reset** through a one-time magic link
- ⚡ **Upstash Redis** for fast, expiring storage (OTPs, reset tokens, sessions)
- 🐘 **PostgreSQL** as the durable source of truth for users
- 🌐 **Header-based auth**: no cookies, no credentialed CORS
- 📐 **Uniform JSON response envelope** across every endpoint
- 🐳 **Dockerized** and ready to deploy anywhere

---

## 🧰 Tech Stack

| Layer            | Technology        | Purpose                                   |
| ---------------- | ----------------- | ----------------------------------------- |
| Language         | **Go**            | Fast, compiled, minimal-footprint service |
| Database         | **PostgreSQL**    | Persistent user storage                   |
| Cache / Store    | **Upstash Redis** | OTPs, reset tokens, token state with TTL  |
| Email            | **Resend**        | Transactional emails (OTP, magic link)    |
| Password hashing | **bcrypt**        | Salted, adaptive password hashing         |
| Tokens           | **JWT**           | Access & refresh tokens                   |
| Packaging        | **Docker**        | Consistent builds and deployments         |

---

## 🗺️ API Overview

Base path: `/api/v1/auth`

| Method | Endpoint           | Auth Required           | Description                          |
| :----: | ------------------ | :---------------------: | ------------------------------------ |
| `POST` | `/signup`          |           ❌            | Create a new account                 |
| `POST` | `/signin`          |           ❌            | Sign in and receive tokens           |
| `POST` | `/signout`         |  ✅ Access token        | Sign out the current session         |
| `POST` | `/send-otp`        |           ❌            | Send email verification OTP          |
| `POST` | `/verify-otp`      |           ❌            | Verify email with the 6-digit OTP    |
| `GET`  | `/refresh`         |  🔁 Refresh token       | Get a new access & refresh token     |
| `POST` | `/reset-password`  |           ❌            | Email a password reset magic link    |
| `POST` | `/update-password` |  🎫 Reset token (query) | Set a new password using the token   |
| `GET`  | `/me`              |  ✅ Access token        | Get the current authenticated user   |

---

## 📦 Response Format

Every endpoint returns the same envelope:

```json
{
  "success": true,
  "message": "message from backend",
  "data": {}
}
```

| Field     | Type      | Description                                                              |
| --------- | --------- | ------------------------------------------------------------------------ |
| `success` | `boolean` | `true` on success, `false` on any failure                                |
| `message` | `string`  | Human-readable message from the backend                                  |
| `data`    | `any`     | Route-specific payload on success. `nil` / omitted when `success: false` |

**Failure example**

```json
{
  "success": false,
  "message": "invalid email or password"
}
```

---

## 🔑 Authentication

- **Access token**: send in the `Authorization` header as a Bearer token:

  ```http
  Authorization: Bearer <access_token>
  ```

- **Refresh token**: send in the `X-Refresh-Token` header (only for `/refresh`):

  ```http
  X-Refresh-Token: <refresh_token>
  ```

- **CORS**: credentials are **not allowed**. Cookies are never used; tokens travel only in headers.

---

## 📡 Endpoint Reference

### 🟢 `POST /api/v1/auth/signup`

Create a new user account.

**Request**

```json
{
  "email": "user@example.com",
  "password": "StrongP@ssw0rd"
}
```

**Response**

```json
{
  "success": true,
  "message": "user created successfully",
  "data": "db_user_id"
}
```

---

### 🟢 `POST /api/v1/auth/signin`

Authenticate with email and password.

**Request**

```json
{
  "email": "user@example.com",
  "password": "StrongP@ssw0rd"
}
```

**Response**

```json
{
  "success": true,
  "message": "signed in successfully",
  "data": {
    "access_token": "eyJhbGciOi...",
    "refresh_token": "eyJhbGciOi..."
  }
}
```

---

### 🟢 `POST /api/v1/auth/signout`

Send the access token in the `Authorization` header.

```http
Authorization: Bearer <access_token>
```

> ⚠️ **Frontend note:** after a successful response, you **must discard both** the access and refresh tokens on the client.

---

### 🟢 `POST /api/v1/auth/send-otp`

Send a 6-digit verification OTP to the user's email.

**Request**

```json
{
  "email": "user@example.com"
}
```

**Response** *(no `data`)*

```json
{
  "success": true,
  "message": "otp sent to email"
}
```

---

### 🟢 `POST /api/v1/auth/verify-otp`

Verify the user's email using the OTP.

**Request**

```json
{
  "email": "user@example.com",
  "otp": "123456"
}
```

**Response**

```json
{
  "success": true,
  "message": "email verified successfully"
}
```

---

### 🔵 `GET /api/v1/auth/refresh`

Exchange a refresh token for a fresh token pair.

```http
X-Refresh-Token: <refresh_token>
```

**Response**

```json
{
  "success": true,
  "message": "token refreshed",
  "data": {
    "access_token": "eyJhbGciOi...",
    "refresh_token": "eyJhbGciOi..."
  }
}
```

---

### 🟢 `POST /api/v1/auth/reset-password`

Request a password reset. A magic link is emailed to the user.

**Request**

```json
{
  "email": "user@example.com"
}
```

**Response**

```json
{
  "success": true,
  "message": "reset link sent to email"
}
```

**Magic link format**

```
https://<frontend_url>/reset-password?token=<uuid>
```

> 💡 **Frontend flow:** when the user lands on this page, show a *new password* field, validate the password client-side, then call `POST /api/v1/auth/update-password?token=<uuid>`.

---

### 🟢 `POST /api/v1/auth/update-password?token=<uuid>`

Set a new password using the token from the magic link.

**Request**

```json
{
  "password": "NewStr0ngP@ssw0rd"
}
```

**Response**

```json
{
  "success": true,
  "message": "password updated successfully"
}
```

---

### 🔵 `GET /api/v1/auth/me`

Get the currently authenticated user.

```http
Authorization: Bearer <access_token>
```

**Response**

```json
{
  "success": true,
  "message": "user fetched successfully",
  "data": "user@example.com"
}
```

---

## 🔄 Flows

### Registration & Verification

```mermaid
sequenceDiagram
    participant C as Client
    participant A as Auth Service
    C->>A: POST /signup
    A-->>C: user id
    C->>A: POST /send-otp
    A-->>C: OTP emailed via Resend
    C->>A: POST /verify-otp
    A-->>C: email verified ✅
```

### Session Lifecycle

```mermaid
sequenceDiagram
    participant C as Client
    participant A as Auth Service
    C->>A: POST /signin
    A-->>C: access_token + refresh_token
    C->>A: GET /me (Bearer access_token)
    A-->>C: user email
    C->>A: GET /refresh (X-Refresh-Token)
    A-->>C: new token pair
    C->>A: POST /signout
    A-->>C: success (client discards tokens)
```

### Password Reset

```mermaid
sequenceDiagram
    participant C as Client
    participant A as Auth Service
    participant E as Email (Resend)
    C->>A: POST /reset-password
    A->>E: send magic link
    E-->>C: link with ?token=uuid
    C->>A: POST /update-password?token=uuid
    A-->>C: password updated ✅
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (1.21+ recommended)
- A [PostgreSQL](https://www.postgresql.org/) database
- An [Upstash Redis](https://upstash.com/) instance
- A [Resend](https://resend.com/) API key
- [Docker](https://www.docker.com/) *(optional)*

### Clone & Run

```bash
# 1. Clone the repository
git clone https://github.com/merafiulhoque/auth-service.git
cd auth-service

# 2. Configure environment
cp .env.example .env
# then fill in your values

# 3. Install dependencies
go mod download

# 4. Run the service
go run ./...
```

---

## ⚙️ Environment Variables

Create a `.env` file in the project root. Variable names below are examples, so match them to your project's config.

```env
# Server
PORT=8080
FRONTEND_URL=https://your-frontend.com

# PostgreSQL
DATABASE_URL=postgres://user:password@host:5432/dbname?sslmode=require

# Upstash Redis
UPSTASH_REDIS_URL=rediss://default:password@host:6379

# JWT
JWT_ACCESS_SECRET=change-me
JWT_REFRESH_SECRET=change-me-too

# Resend
RESEND_API_KEY=re_xxxxxxxxx
EMAIL_FROM=noreply@yourdomain.com
```

> 🔒 **Never commit your `.env` file.** Use strong, unique, random secrets for JWT signing.

---

## 🐳 Docker

```bash
# Build the image
docker build -t auth-service .

# Run the container
docker run -p 8080:8080 --env-file .env auth-service
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

1. Fork the repo
2. Create your branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m "Add amazing feature"`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

---

<div align="center">

Built with ❤️ and Go by [**@merafiulhoque**](https://github.com/merafiulhoque)

⭐ If this project helped you, consider giving it a star!

</div>