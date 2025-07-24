
# Task Manager API with Go + MongoDB

This is a simple RESTful API built using **Golang**, **Gin**, and **MongoDB** to manage tasks. It supports basic CRUD operations like creating, reading, updating, and deleting tasks.

---

## Features

- Add a new task
- Retrieve all or specific task by ID
- Update existing task
- Delete task
- Connects to MongoDB Atlas with graceful shutdown

---

## Prerequisites

- ✅ Go 1.20+
- ✅ MongoDB Atlas account with a valid connection URI
- ✅ Internet connection to connect to your MongoDB cluster

---

## MongoDB Configuration

1. Create a **MongoDB Atlas cluster**
2. Whitelist your IP address or allow all with `0.0.0.0/0`
3. Replace the connection string in `data/database.go`:

```go
opts := options.Client().ApplyURI("mongodb+srv://<username>:<password>@<cluster-url>/?retryWrites=true&w=majority")
```

4. Database: `Task_Database`  
   Collection: `Task_Five`

---

## Project Structure

```plaintext
task_manager_api_mongodb/
├── main.go                    → Application entry point
├── controllers/              → API handlers for routes
│   └── task_controller.go
├── data/                     → MongoDB config and operations
│   ├── database.go
│   └── task_service.go
├── model/                    → Task struct model
│   └── task.go
├── routes/                   → Route grouping using Gin
│   └── task_routes.go
```

---

## Running the API

```bash
go run main.go
```

Once the server starts, it will be available at:  
`http://localhost:8080`

---

## API Endpoints

### Create Task

- **POST** `/tasks`
- **Request Body**:

```json
{
  "id": "1",
  "title": "Learn Go",
  "description": "Practice with Gin and MongoDB"
}
```

- **Response**:

```json
{
  "message": "Task Added"
}
```

---

### Get All Tasks

- **GET** `/tasks`

- **Response**:

```json
[
  {
    "id": "1",
    "title": "Learn Go",
    "description": "Practice with Gin and MongoDB"
  }
]
```

---

### Get Task by ID

- **GET** `/tasks/:id`
- **URL Param**: `id` (integer)

- **Response**:

```json
{
  "id": "1",
  "title": "Learn Go",
  "description": "Practice with Gin and MongoDB"
}
```

---

### Update Task

- **PUT** `/tasks/:id`
- **Request Body**:

```json
{
  "title": "Updated Title",
  "description": "Updated Description"
}
```

- **Response**:

```json
{
  "title": "Updated Title",
  "description": "Updated Description"
}
```

---

### Delete Task

- **DELETE** `/tasks/:id`

- **Response**:

```json
{
  "message": "Successfully deleted task"
}
```

---
