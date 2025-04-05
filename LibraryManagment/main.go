package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"LibraryManagment/controllers"
	"LibraryManagment/services"
)

func main() {
	service := services.NewLibraryServices()
	controller := controllers.NewLibraryController(service)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Enter Book ID: ")
			id := readInt(reader)
			fmt.Print("Enter Book Title: ")
			title := readString(reader)
			fmt.Print("Enter Book Author: ")
			author := readString(reader)
			controller.AddBook(id, title, author)

		case 2:
			fmt.Print("Enter Book ID to remove: ")
			id := readInt(reader)
			controller.RemoveBook(id)

		case 3:
			fmt.Print("Enter Book ID to borrow: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			controller.BorrowBook(bookID, memberID)

		case 4:
			fmt.Print("Enter Book ID to return: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			controller.ReturnBook(bookID, memberID)

		case 5:
			controller.ListAvailableBooks()

		case 6:
			controller.ListBorrowedBooks()

		case 0:
			fmt.Println("Exiting Library System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Utility Functions
func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number. Try again:")
		return readInt(reader)
	}
	return val
}

func readString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
