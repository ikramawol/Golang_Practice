package controllers

import (
	"LibraryManagment/models"
	"LibraryManagment/services"
	"fmt"
)

type LibraryController struct {
	Service *services.LibraryServices
}

func NewLibraryController(service *services.LibraryServices) *LibraryController {
	return &LibraryController{Service: service}
}

func (lc *LibraryController) AddBook(id int, author, title string) {
	book := models.Book{ID: id, Title: title, Author: author}
	lc.Service.AddBook(book)
	fmt.Println("Book added successfully:", book)
}

func (lc *LibraryController) RemoveBook(id int) {
	lc.Service.RemoveBook(id)
	fmt.Println("Book removed successfully")
}

func (lc *LibraryController) BorrowBook(id int, memeberId int) {
	lc.Service.BorrowBook(id, memeberId)
	fmt.Println("Book borrowed successfully")
}

func (lc *LibraryController) ReturnBook(id int, memberId int) {
	lc.Service.ReturnBook(id, memberId)
	fmt.Println("Book returned successfully")
}

func (lc *LibraryController) ListAvailableBooks() {

	books := lc.Service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available in the library.")
		return
	}
	fmt.Println("Available books in the library:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Auther: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) ListBorrowedBooks() {

	books := lc.Service.ListBorrowedBooks()
	if len(books) == 0 {
		fmt.Println("No books are currently borrowed.")
		return
	}
	fmt.Println("Currently borrowed books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
