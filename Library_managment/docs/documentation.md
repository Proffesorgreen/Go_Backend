# Library Management System

## Overview

This is a basic console-based application designed to manage books and members in a library. It allows you to perform core operations like adding, removing, borrowing, and returning books, and listing available or borrowed books.

## Project Structure

The project is organized into clear directories to separate different parts of the application:

* `models/`: Contains the data structures.
* `services/`: Holds the main business logic (how books are added, removed, borrowed, etc.). This is where the core library operations happen.
* `controllers/`: Acts as the bridge between the user interface (the console menu) and the `services`. It handles user input and calls the right service functions.
* `main.go`: The starting point of your application. It sets up the library, the controller, and runs the main menu loop.

## How to Run

To get the library system up and running:

1.  **Save the files:** Ensure all your Go files (`.go`) are saved in their respective directories within the `library_management` folder.
    * `library_management/models/book.go`
    * `library_management/models/member.go`
    * `library_management/services/library_service.go`
    * `library_management/controllers/library_controller.go`
    * `library_management/main.go`
2.  **Open your terminal** or command prompt.
3.  **Navigate to the `library_management` directory** (where your `main.go` file is located).

    ```bash
    cd path/to/your/library_management
    ```
4.  **Run the application** using the Go command:

    ```bash
    go run main.go
    ```
5.  The system will start, display a welcome message, initialize some sample data, and present the main menu.

## 📖 Key Components Explained

### Models

* **`Book`**: Represents a book with an **ID**, **Title**, **Author**, and **Status_book** (which can be "Available" or "Borrowed").
* **`Memeber`**: Represents a library member with an **ID**, **Name**, and a list of **Borrowed_Books**.

### Services

The `services` package contains the core logic for managing the library. The `Library` struct holds all the **Books_ID** (all books in the library) and **Memebers_ID** (all registered members).

* **`AddBook(book models.Book)`**: Adds a new book to the library and assigns it a unique ID.
* **`RemoveBook(bookID int)`**: Removes a book by its ID.
* **`BorrowBook(bookID int, memberID int) error`**: Handles the process of a member borrowing a book.
* **`ReturnBook(bookID int, memberID int) error`**: Handles a member returning a book.
* **`ListAvailableBooks() []models.Book`**: Returns a list of all books currently available.
* **`ListBorrowedBooks(memberID int) []models.Book`**: Returns a list of books borrowed by a specific member.
* **`getbookbyid(bookID int) *models.Book`**: An internal helper to find a book by its ID.

### Controllers

The `LibraryController` acts as the intermediary. It takes user input from the console, processes it, and then calls the appropriate method on your `services.Library` instance. It ensures that all actions are performed on the same library data throughout the application's runtime.

### Main Application (`main.go`)

This is where everything comes together.

* It creates a **single instance** of your `services.Library`. This is crucial because it ensures that all operations (adding books, borrowing, etc.) are performed on the same set of data, so your changes persist throughout your session.
* It initializes the `LibraryController`, passing it a reference to that single `Library` instance.
* It sets up some initial sample books and one member (ID: 1, Name: "Test Member") so you have data to work with immediately.
* It runs the main loop that continuously displays the menu and handles your choices.
