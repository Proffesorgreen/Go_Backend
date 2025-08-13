# Task Manager API Documentation

## Introduction

This document provides a comprehensive overview of the Task Manager API, its architecture, and guidelines for future development. The API is built using Go and follows the principles of Clean Architecture to ensure a modular, scalable, and maintainable codebase.

## Architecture

The application is structured based on the Clean Architecture pattern, which separates concerns into distinct layers. This design promotes loose coupling and high cohesion, making the system easier to understand, test, and evolve.

The main layers of the application are:

-   **Domain:** The core of the application, containing the business logic and entities.
-   **Repository:** The data access layer, responsible for interacting with the database.
-   **Use Cases:** The application-specific business rules, orchestrating the flow of data between the domain and the repositories.
-   **Delivery:** The presentation layer, responsible for handling user requests and responses.

### Layers

#### Domain

The `domain` layer is the heart of the application. It contains the core business logic and entities, such as `Task` and `User`. This layer is independent of any external frameworks or libraries, ensuring that the business rules can be tested in isolation.

#### Repository

The `repository` layer is responsible for all data-related operations. It defines interfaces for accessing the database, which are then implemented by the `repository` package. This abstraction allows the underlying data source to be swapped out with minimal impact on the rest of the application.

#### Use Cases

The `usecases` layer contains the application-specific business logic. It orchestrates the flow of data between the `domain` and `repository` layers to perform specific tasks, such as creating a new task or registering a new user.

#### Delivery

The `delivery` layer is the entry point to the application. It handles incoming HTTP requests, validates user input, and calls the appropriate use cases to perform the requested actions. This layer is responsible for presenting the data to the user in a clear and concise format.

## API Endpoints

The following is a list of all available API endpoints and their functionalities:

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | /user/register | Registers a new user. |
| POST | /user/login | Logs in an existing user. |
| POST | /tasks/ | Creates a new task. |
| PATCH | /tasks/:id | Updates an existing task. |
| DELETE | /tasks/:id | Deletes an existing task. |
| GET | /tasks/ | Retrieves all tasks. |
| GET | /tasks/:id | Retrieves a single task by its ID. |

## Unit Test Suite

The project includes a comprehensive unit test suite to ensure the correctness and reliability of the codebase.

### Running Tests

To run all unit tests in the project, navigate to the project's root directory and execute the following command:

```bash
go test ./...
```

This command will discover and run all `_test.go` files within the current directory and its subdirectories.

### Test Coverage

Test coverage measures the percentage of your codebase that is executed by your tests. High test coverage indicates that a larger portion of your code is being tested, which can help in identifying untested areas and potential bugs.

To generate a test coverage report, use the following commands:

1.  **Generate coverage data:**
    ```bash
    go test -coverprofile=cover.out ./...
    ```
    This command runs the tests and writes coverage profiles to a file named `cover.out`.

2.  **View coverage report in HTML:**
    ```bash
    go tool cover -html=cover.out
    ```
    This command opens an HTML report in your web browser, visually indicating which parts of your code are covered by tests and which are not.
