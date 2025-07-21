package services

import (
	"fmt"
	"library_managment/models"
)

type Library_Manager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Memebers_ID []models.Memeber
	Books_ID    []models.Book
}

var next_Book_ID = 1

func (l *Library) AddBook(book models.Book) {
	book.ID = next_Book_ID
	book.Status_book = "Available"
	l.Books_ID = append(l.Books_ID, book)
	next_Book_ID += 1
	fmt.Print("Sucessfully added the book to the library")
}

func (l *Library) RemoveBook(bookID int) {
	for idx := range l.Books_ID {
		if l.Books_ID[idx].ID == bookID {
			l.Books_ID = append(l.Books_ID[:idx], l.Books_ID[idx+1:]...)
			fmt.Printf("Sucessfully removed book %d from the library", bookID)
			return
		}
	}
	fmt.Printf("There is no book with the ID %d in the Library", bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	for idx := range l.Memebers_ID {
		if l.Memebers_ID[idx].ID == memberID {
			if bookID > 0 && bookID <= len(l.Books_ID) {
				book_to_borrow := l.getbookbyid(bookID)
				if book_to_borrow.Status_book == "Available" {
					book_to_borrow.Status_book = "Borrowed"
					l.Memebers_ID[idx].Borrowed_Books = append(l.Memebers_ID[idx].Borrowed_Books, *book_to_borrow)
					fmt.Print("Sucessfully Borrowed the book from the library")
					return nil
				} else {
					fmt.Print("The Book is already Borrowed by Others")
					return nil
				}

			} else {
				fmt.Print("Invalid Book ID")
				return nil
			}
		}
	}
	fmt.Print("Invalid Member ID")
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	for idx := range l.Memebers_ID {
		if l.Memebers_ID[idx].ID == memberID {
			if bookID > 0 && bookID <= len(l.Books_ID) {
				for jdx, book := range l.Memebers_ID[idx].Borrowed_Books {
					if book.ID == bookID {
						book_to_borrow := l.getbookbyid(bookID)
						book_to_borrow.Status_book = "Available"
						l.Memebers_ID[idx].Borrowed_Books = append(l.Memebers_ID[idx].Borrowed_Books[:jdx], l.Memebers_ID[idx].Borrowed_Books[jdx+1:]...)
						return nil
					}
				}
			} else {
				fmt.Print("Invalid Book ID")
				return nil
			}
		}
	}
	fmt.Print("Invalid Member ID")
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var avaliable_books []models.Book
	for _, book := range l.Books_ID {
		if book.Status_book == "Available" {
			avaliable_books = append(avaliable_books, book)
		}
	}
	return avaliable_books
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	for _, memeber := range l.Memebers_ID {
		if memeber.ID == memberID {
			return memeber.Borrowed_Books
		}
	}
	fmt.Print("Invalid Memeber ID")
	return []models.Book{}
}

func (l *Library) getbookbyid(bookID int) *models.Book {
	for idx := range l.Books_ID {
		if l.Books_ID[idx].ID == bookID {
			return &l.Books_ID[idx]
		}
	}
	return nil
}