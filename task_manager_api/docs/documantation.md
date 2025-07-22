# Task Manager API

## Overview

This is a simple RESTful API built using **Golang** and the **Gin web framework**. It allows you to manage a collection of tasks — create, read, update, and delete (CRUD) — using HTTP requests.

---

## Project Structure

The project is organized for clarity and maintainability:

- `models/`: Defines the structure of a **Task** (fields like ID, Title, Description, etc.).
- `controllers/`: Handles incoming HTTP requests and routes them to appropriate functions.
- `data/` or `services/`: Contains core logic and operations on task data (in-memory or persistent).
- `router/`: Configures all the available API routes using the Gin framework.
- `main.go`: Entry point of the application. Initializes the server and starts listening for requests.

---

## How to Run

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/task_manager_api.git
    cd task_manager_api
    ```

2. Install dependencies and tidy up modules:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

4. The server will start on:

    ```
    http://localhost:8080
    ```

---

## 🔗 API Endpoints

| Method | Endpoint        | Description                 |
|--------|------------------|-----------------------------|
| GET    | `/tasks`         | Get all tasks               |
| GET    | `/tasks/:id`     | Get a specific task by ID   |
| POST   | `/tasks`         | Create a new task           |
| PUT    | `/tasks/:id`     | Update an existing task     |
| DELETE | `/tasks/:id`     | Delete a task by ID         |

---
