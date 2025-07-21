package main

import (
	"fmt"
	controllers "library_managment/controller"
	"library_managment/models"
	"library_managment/services"
)

func main() {
	fmt.Println("Welcome to the library management system")
	library := &services.Library{
		Books_ID:    []models.Book{},
		Memebers_ID: []models.Memeber{},
	}

	controller := controllers.NewLibraryController(library)
	for {
		var input int

		fmt.Println("\nMenu")
		fmt.Println("------------------")
		fmt.Println("1. List all available books")
		fmt.Println("2. List all borrowed books by a member")
		fmt.Println("3. Add a book")
		fmt.Println("4. Remove a book")
		fmt.Println("5. Borrow a book")
		fmt.Println("6. Return a book")
		fmt.Println("-1. Exit Program")
		fmt.Println("------------------")
		fmt.Println("Enter your choice: ")

		fmt.Scanf("%d\n", &input)
		switch input {
		case -1:
			fmt.Println("Exiting Program")
			return
		case 1:
			controller.HandleListAvailableBooks()
		case 2:
			controller.HandleListBorrowedBooks()
		case 3:
			controller.HandleAddBook()
		case 4:
			controller.HandleRemoveBook()
		case 5:
			controller.HandleBorrowBook()
		case 6:
			controller.HandleReturnBook()
		default:
			continue
		}
	}
}
