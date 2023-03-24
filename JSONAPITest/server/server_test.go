package server

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddBook(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"
	bAdded, mBookList, err := AddBook(ID, Title, Author, PublicationDate)
	if err != nil {
		t.Errorf("Error returned %s\n", err.Error())
	}
	assert.Equal(t, true, bAdded)
	assert.Equal(t, 1, len(mBookList))
	assert.Equal(t, ID, inmem_booklist[0].ID)
	assert.Equal(t, Title, inmem_booklist[0].Title)
	assert.Equal(t, Author, inmem_booklist[0].Author)
	assert.Equal(t, PublicationDate, inmem_booklist[0].PublicationDate)
}

// test add book - missing elements... table driven tests...
func Test_AddBook_BadSubmission(t *testing.T) {
	testcases := []struct {
		ID                  string
		Title               string
		Author              string
		PublicationDate     string
		ReturnedErrorString string
		ReturnedBoolean     bool
	}{
		{"", "Hello World", "JKRowling", "01/02/2010", "Missing ID", false},
		{"2", "", "JKRowling", "01/02/2010", "Missing Title", false},
		{"3", "Hello World", "", "01/02/2010", "Missing Author", false},
		{"4", "Hello World", "JKRowling", "", "Missing PublicationDate", false},
	}
	for _, testcase := range testcases {
		bAdded, _, err := AddBook(testcase.ID, testcase.Title, testcase.Author, testcase.PublicationDate)
		assert.Equal(t, bAdded, testcase.ReturnedBoolean)
		assert.Equal(t, err.Error(), testcase.ReturnedErrorString)

	}
}

// haven't yet checked for supplying duplicated book ID..
func Test_DuplicateBookID(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"
	bAdded, _, _ := AddBook(ID, Title, Author, PublicationDate)
	assert.Equal(t, true, bAdded)
	bAdded, _, err := AddBook(ID, Title, Author, PublicationDate)
	assert.Equal(t, false, bAdded)
	assert.Equal(t, err.Error(), "Duplicate ID")

}

// haven't yet checked for badly formatted publication dates..

func Test_GetListAllBooks(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"
	bAdded, mBookList, err := AddBook(ID, Title, Author, PublicationDate)
	if err != nil {
		t.Errorf("Error returned %s", err.Error())
	}
	assert.Equal(t, true, bAdded)
	mCompleteList := GetBookList()
	if !reflect.DeepEqual(mBookList, mCompleteList) {
		t.Errorf("Didn't get complete list")
	}

}

func Test_DeleteBook(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"

	testcases := []struct {
		StartDeleteBookListLength int
		ID                        string
		PostDeleteBookListLength  int
		HasBeenDeleted            bool
	}{
		{1, "1", 0, true},
		{1, "2", 1, false},
	}
	for _, testcase := range testcases {
		bAdded, mBookList, err := AddBook(ID, Title, Author, PublicationDate)
		if err != nil {
			t.Errorf("Error returned %s", err.Error())
		}
		assert.Equal(t, true, bAdded)
		assert.Equal(t, testcase.StartDeleteBookListLength, len(mBookList))
		bDeleted, mNewBookList, err := DeleteBook(testcase.ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, testcase.HasBeenDeleted, bDeleted)
		assert.Equal(t, testcase.PostDeleteBookListLength, len(mNewBookList))
	}

}

func Test_GetBookByID(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"

	testcases := []struct {
		ID         string
		TitleFound string
		FoundBook  bool
	}{
		{"1", "Revised Standard Version Bible", true},
		{"2", "", false},
	}
	bAdded, _, err := AddBook(ID, Title, Author, PublicationDate)
	if err != nil {
		t.Errorf("Error returned %s", err.Error())
	}
	assert.Equal(t, true, bAdded)
	for _, testcase := range testcases {
		bFoundBook, mBook, err := GetBookByID(testcase.ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, testcase.FoundBook, bFoundBook)
		assert.Equal(t, testcase.TitleFound, mBook.Title)
	}

}

func Test_UpdateBook(t *testing.T) {
	var ID, Title, Author, PublicationDate string
	ID = "1"
	Title = "Revised Standard Version Bible"
	Author = "Unknown"
	PublicationDate = "31/12/1980"

	testcases := []struct {
		ID                      string
		NewTitle                string
		NewAuthor               string
		NewPublicationDate      string
		ExpectedTitle           string
		ExpectedAuthor          string
		ExpectedPublicationDate string
		UpdatedBook             bool
	}{
		// note th booleans aren't in the API - there for test purposes
		{"1", "", "Various", "", "Revised Standard Version Bible", "Various", "31/12/1980", true},
		{"1", "Revised Standard Version Holy Bible", "", "", "Revised Standard Version Holy Bible", "Various", "31/12/1980", true},
		{"2", "Revised Standard Version Holy Bible", "", "", "", "", "", false},
		{"1", "", "", "01/01/2010", "Revised Standard Version Holy Bible", "Various", "01/01/2010", true},
	}
	bAdded, _, err := AddBook(ID, Title, Author, PublicationDate)
	if err != nil {
		t.Errorf("Error returned %s", err.Error())
	}
	assert.Equal(t, true, bAdded)
	for _, testcase := range testcases {
		bUpdatedBook, mBook, err := UpdateBook(testcase.ID, testcase.NewTitle, testcase.NewAuthor, testcase.NewPublicationDate)
		assert.Equal(t, nil, err)
		assert.Equal(t, testcase.UpdatedBook, bUpdatedBook)
		assert.Equal(t, testcase.ExpectedTitle, mBook.Title)
		assert.Equal(t, testcase.ExpectedAuthor, mBook.Author)
		assert.Equal(t, testcase.ExpectedPublicationDate, mBook.PublicationDate)
	}

}
