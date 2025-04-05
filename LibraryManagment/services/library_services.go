package services

import (
	"LibraryManagment/models"

	"errors"
	"fmt"
)

type LibraryServices struct {
	books   map[int]models.Book
	members map[int][]int
}

func NewLibraryServices() *LibraryServices {
	return &LibraryServices{
		books:   make(map[int]models.Book),
		members: make(map[int][]int),
	}
}

func (ls *LibraryServices) AddBook(book models.Book) {
	if _, exists := ls.books[book.ID]; exists {
		fmt.Println("Book already exists")
		return
	}
	ls.books[book.ID] = book
}

func (ls *LibraryServices) RemoveBook(id int) {
	if _, exists := ls.books[id]; !exists {
		fmt.Println("Book does not exixt")
		return
	}
	delete(ls.books, id)
}

func (ls *LibraryServices) BorrowBook(id int, memberId int) error {
	book, exists := ls.books[id]
	if !exists {
		return errors.New("Book not found!")
	}
	if book.IsBorrowed {
		return errors.New("Book is already borrowed")
	}
	book.IsBorrowed = true
	ls.books[id] = book
	ls.members[memberId] = append(ls.members[memberId], id)
	return nil
}

func (ls *LibraryServices) ReturnBook(id int, memberId int) error {
	book, exists := ls.books[id]
	if !exists {
		return errors.New("Book not found!")
	}
	if !book.IsBorrowed {
		return errors.New("Book is not borrowed")
	}

	book.IsBorrowed = false
	ls.books[id] = book
	return nil

}

func (ls *LibraryServices) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range ls.books {
		if !book.IsBorrowed {
			availableBooks = append(availableBooks, book) // add to available books list if not borrowed
		}
	}
	return availableBooks
}

func (ls *LibraryServices) ListBorrowedBooks() []models.Book {
	var borrowedBooks []models.Book
	for _, book := range ls.books {
		if book.IsBorrowed {
			borrowedBooks = append(borrowedBooks, book) // add to borrowed books list if borrowed
		}
	}
	if len(borrowedBooks) == 0 {
		return []models.Book{} // return empty slice if no borrowed books found
	}
	return borrowedBooks
}
