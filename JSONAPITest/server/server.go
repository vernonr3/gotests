// This server allows a user to
// Get a list of all books
// Get a single book by ID
// Add a new book
// Update an existing book
// Delete a book
//
// "encoding/json"
package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Book is described by the following properties
type Book struct {
	ID              string `json:"ID"`
	Title           string `json:"Title"`
	Author          string `json:"Author"`
	PublicationDate string `json:"PublicationDate"`
}

type BookList []Book

var inmem_booklist BookList

// AddBook enables a book to be added to the list
// duplicates are determined by matching ID.
// if 2 books have the different IDs any of the other fields can match without a problem
// All Fields have to be submitted in order for a book to be correctly added
// failing to add a book because of duplication will be treated as an error submission
func AddBook(ID string, Title string, Author string, PublicationDate string) (bool, BookList, error) {
	var newBook Book
	for i := 0; i < len(inmem_booklist); i++ {
		if inmem_booklist[i].ID == ID {
			return false, inmem_booklist, errors.New("Duplicate ID")
		}
	}
	if len(ID) == 0 {
		return false, inmem_booklist, errors.New("Missing ID")
	}
	if len(Title) == 0 {
		return false, inmem_booklist, errors.New("Missing Title")
	}
	if len(Author) == 0 {
		return false, inmem_booklist, errors.New("Missing Author")
	}
	if len(PublicationDate) == 0 {
		return false, inmem_booklist, errors.New("Missing PublicationDate")
	}

	newBook.ID = ID
	newBook.Title = Title
	newBook.Author = Author
	newBook.PublicationDate = PublicationDate
	fmt.Printf("Added %v\n", newBook)
	inmem_booklist = append(inmem_booklist, newBook)
	return true, inmem_booklist, nil
}

// Delete a book
// The deletion will be based on the ID of the book
// no other fields will be checked
func DeleteBook(ID string) (bool, BookList, error) {
	var bFoundID bool
	for i := 0; i < len(inmem_booklist); i++ {
		if inmem_booklist[i].ID == ID {
			// shuffle the remaining items up to close the space..
			inmem_booklist = append(inmem_booklist[:i], inmem_booklist[i+1:]...)
			bFoundID = true
		}
	}
	return bFoundID, inmem_booklist, nil
}

func GetBookByID(ID string) (bool, Book, error) {
	var bFoundID bool
	var bookToFind Book
	for i := 0; i < len(inmem_booklist); i++ {
		if inmem_booklist[i].ID == ID {
			// shuffle the remaining items up to close the space..
			bookToFind = inmem_booklist[i]
			bFoundID = true
		}
	}
	return bFoundID, bookToFind, nil
}

func UpdateBook(ID string, NewTitle string, NewAuthor string, NewPublicationDate string) (bool, Book, error) {
	var bFoundID bool
	var bookToFind Book
	for i := 0; i < len(inmem_booklist); i++ {
		if inmem_booklist[i].ID == ID {
			// shuffle the remaining items up to close the space..
			if len(NewTitle) > 0 {
				inmem_booklist[i].Title = NewTitle
			}
			if len(NewAuthor) > 0 {
				inmem_booklist[i].Author = NewAuthor
			}
			if len(NewPublicationDate) > 0 {
				inmem_booklist[i].PublicationDate = NewPublicationDate
			}
			bookToFind = inmem_booklist[i]
			bFoundID = true
		}
	}
	return bFoundID, bookToFind, nil
}

// Get the complete booklist
func GetBookList() BookList {
	return inmem_booklist
}

// Start the Server with the multiplexer etc.
func Start() {
	fmt.Println("Server alive on port 8010")
	r := mux.NewRouter()
	r.HandleFunc("/AddBook", AddBookToList)
	r.HandleFunc("/GetBookList", GetCompleteBookList)
	r.HandleFunc("/GetBookByID", GetABookByID)
	r.HandleFunc("/DeleteBook", DeleteABook)
	r.HandleFunc("/UpdateBook", UpdateABook)
	http.ListenAndServe("localhost:8010", r)
}
