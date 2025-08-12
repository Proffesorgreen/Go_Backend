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

## Future Development

To add new features to the application, follow these guidelines:

1.  **Define the domain:** Start by defining the new entities and business logic in the `domain` layer.
2.  **Implement the repository:** Create a new repository in the `repository` layer to handle data access for the new entities.
3.  **Create the use cases:** Implement the new use cases in the `usecases` layer to orchestrate the flow of data.
4.  **Add the delivery handlers:** Create new handlers in the `delivery` layer to expose the new functionality through the API.

By following these guidelines, you can ensure that the new features are well-integrated into the existing architecture and that the codebase remains clean and maintainable.
