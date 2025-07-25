# Task Manager API

A RESTful API for managing tasks with user authentication and role-based access control, built with Go (Gin framework) and MongoDB.

---

## Features

- User registration and login with JWT authentication
- Role-based authorization (admin vs regular user)
- CRUD operations for tasks
- Secure password hashing with bcrypt
- Graceful shutdown and MongoDB connection management

---

## Technologies Used

- Go (Golang)
- Gin Web Framework
- MongoDB
- JWT for authentication
- bcrypt for password hashing

---

## Getting Started

### Prerequisites

- Go 1.18+
- MongoDB instance running
- Set environment variable `JWT_SECRET` (used for signing JWT tokens)
- Optional: Set `PORT` environment variable (defaults to 8080)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/task_manager_api.git
cd task_manager_api
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set environment variables:

```bash
export JWT_SECRET=your_secret_key
export PORT=8080 # optional
```

4. Run the server:

```bash
go run main.go
```

---

## API Endpoints

### Authentication

| Method | Endpoint     | Description             | Auth Required | Role Required |
|--------|--------------|-------------------------|---------------|---------------|
| POST   | `/register`  | Register a new user      | No            | N/A           |
| POST   | `/login`     | Login and get JWT token  | No            | N/A           |

### Tasks

| Method | Endpoint         | Description                 | Auth Required | Role Required |
|--------|------------------|-----------------------------|---------------|---------------|
| GET    | `/tasks`         | Get all tasks               | Yes           | Any           |
| GET    | `/tasks/:id`     | Get task by ID             | Yes           | Any           |
| POST   | `/tasks`         | Create a new task           | Yes           | Admin         |
| PUT    | `/tasks/:id`     | Update a task               | Yes           | Admin         |
| DELETE | `/tasks/:id`     | Delete a task               | Yes           | Admin         |

---

## Project Structure

```
task_manager_api/
├── controllers/       # Handlers for HTTP requests
├── data/              # Database interaction and business logic
├── middleware/        # Authentication and authorization middleware
├── model/             # Data models (User, Task, etc.)
├── routes/            # API route definitions
├── main.go            # Application entry point
├── go.mod             # Go modules file
└── go.sum             # Go modules checksum
```

---

