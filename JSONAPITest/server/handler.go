package server

import (
	"encoding/json"
	"net/http"
)

type ErrorCode struct {
	Error string `json:"error"`
}

// Handle URL /AddBook
func AddBookToList(w http.ResponseWriter, r *http.Request) {
	var mBook Book
	// decode the URL parameters or JSON
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mBook)
	if err != nil {
		panic(err)
	}
	bAdded, mBookList, err := AddBook(mBook.ID, mBook.Title, mBook.Author, mBook.PublicationDate)
	if bAdded && mBookList != nil {
		jsondata, err := json.Marshal(mBookList)
		if err != nil {
			panic("Failed to parse booklist data ")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
		return
	}
	if !bAdded {
		var mError ErrorCode
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		mError.Error = err.Error()
		jsondata, err := json.Marshal(mError)
		if err != nil {
			panic("Failed to parse error data ")
		}
		w.Write(jsondata)
	}
}

// Handle URL /GetBookList
func GetCompleteBookList(w http.ResponseWriter, r *http.Request) {
	mBookList := GetBookList()
	if mBookList != nil {
		jsondata, err := json.Marshal(mBookList)
		if err != nil {
			panic("Failed to parse booklist data ")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
		return
	} else {
		var mError ErrorCode
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		mError.Error = "Failed to retrieve book list"
		jsondata, err := json.Marshal(mError)
		if err != nil {
			panic("Failed to parse error data ")
		}
		w.Write(jsondata)
	}
}

// Handle URL /GetBookByID
func GetABookByID(w http.ResponseWriter, r *http.Request) {
	var mBook Book
	// decode the URL parameters or JSON
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mBook)
	if err != nil {
		panic(err)
	}
	bFound, mBook, err := GetBookByID(mBook.ID)
	if bFound {
		jsondata, err := json.Marshal(mBook)
		if err != nil {
			panic("Failed to parse booklist data ")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
		return
	}
	if !bFound {
		var mError ErrorCode
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		mError.Error = err.Error()
		jsondata, err := json.Marshal(mError)
		if err != nil {
			panic("Failed to parse error data ")
		}
		w.Write(jsondata)
	}
}

// Handle /DeleteBook
func DeleteABook(w http.ResponseWriter, r *http.Request) {
	var mBook Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mBook)
	if err != nil {
		panic(err)
	}
	bDeleted, mBookList, err := DeleteBook(mBook.ID)
	if bDeleted {
		jsondata, err := json.Marshal(mBookList)
		if err != nil {
			panic("Failed to parse booklist data ")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
	}
	return

}

// handle URL /UpdateBook
func UpdateABook(w http.ResponseWriter, r *http.Request) {
	var mBook Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mBook)
	if err != nil {
		panic(err)
	}
	bUpdated, mBookList, err := UpdateBook(mBook.ID, mBook.Title, mBook.Author, mBook.PublicationDate)
	if bUpdated {
		jsondata, err := json.Marshal(mBookList)
		if err != nil {
			panic("Failed to parse booklist data ")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
	}
	if !bUpdated {
		var mError ErrorCode
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		mError.Error = err.Error()
		jsondata, err := json.Marshal(mError)
		if err != nil {
			panic("Failed to parse error data ")
		}
		w.Write(jsondata)
	}
	return

}
