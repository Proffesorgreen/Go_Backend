package controllers

import (
	"bufio"
	"fmt"
	"library_managment/models"
	"library_managment/services" // Import your services package
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	library *services.Library 
	reader  *bufio.Reader   
}

func NewLibraryController(lib *services.Library) *LibraryController {
	return &LibraryController{
		library: lib,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (lc *LibraryController) readString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := lc.reader.ReadString('\n') 
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

func (lc *LibraryController) readInt(prompt string) (int, error) {
	inputStr, err := lc.readString(prompt)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(inputStr) // Converts string to integer
	if err != nil {
		return 0, fmt.Errorf("invalid number input. Please enter a valid integer: %w", err)
	}
	return num, nil
}

func (lc *LibraryController) HandleAddBook() {
	title, err := lc.readString("Enter book title: ")
	if err != nil || title == "" {
		fmt.Println("Error: Invalid title. Title cannot be empty.")
		return
	}

	author, err := lc.readString("Enter book author: ")
	if err != nil || author == "" {
		fmt.Println("Error: Invalid author. Author cannot be empty.")
		return
	}

	book := models.Book{
		Title:  title,
		Author: author,
	}

	lc.library.AddBook(book) 
	fmt.Println()            
}

func (lc *LibraryController) HandleRemoveBook() {
	bookID, err := lc.readInt("Enter the ID of the book to remove: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	lc.library.RemoveBook(bookID) 
	fmt.Println()               
}

func (lc *LibraryController) HandleBorrowBook() {
	bookID, err := lc.readInt("Enter the ID of the book to borrow: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	memberID, err := lc.readInt("Enter the ID of the member borrowing: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	lc.library.BorrowBook(bookID, memberID)
	fmt.Println()
}

func (lc *LibraryController) HandleReturnBook() {
	bookID, err := lc.readInt("Enter the ID of the book to return: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	memberID, err := lc.readInt("Enter the ID of the member returning: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	lc.library.ReturnBook(bookID, memberID)
	fmt.Println()
}

func (lc *LibraryController) HandleListAvailableBooks() {
	books := lc.library.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No books are currently available in the library.")
		return
	}

	fmt.Println("\n--- Available Books ---")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: '%s', Author: '%s', Status: %s\n", book.ID, book.Title, book.Author, book.Status_book)
	}
	fmt.Println("-----------------------")
}

func (lc *LibraryController) HandleListBorrowedBooks() {
	memberID, err := lc.readInt("Enter the ID of the member to list borrowed books for: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	borrowedBooks := lc.library.ListBorrowedBooks(memberID)
	fmt.Println()

	if len(borrowedBooks) == 0 && memberID != 0 {
		fmt.Printf("Member ID %d has not borrowed any books.\n", memberID)
		return
	}

	if len(borrowedBooks) > 0 {
		fmt.Printf("\n--- Books borrowed by Member ID %d ---\n", memberID)
		for _, book := range borrowedBooks {
			fmt.Printf("ID: %d, Title: '%s', Author: '%s'\n", book.ID, book.Title, book.Author)
		}
		fmt.Println("------------------------------------")
	}
}
