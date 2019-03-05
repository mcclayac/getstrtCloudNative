package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Book type with Name Author and ISBN
type Book struct {
	// define book
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	ISBN	string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"0-671-62964-6" : Book{Title:"The Hitchhiker's Guide to the Galaxy", Author:"Douglas Adams",ISBN:"0-671-62964-6"},
	"0123456":Book{Title:"Cloud Native Go", Author:"M.L.Reimer",ISBN:"0123456"},
	"0-345-25855-X":Book{Title:"A Spell for Chameleon", Author:"Piers Anthony",ISBN:"0-345-25855-X"},
}





func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		//books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		book := FromJson(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported method"))
		//fmt.Fprintf(w,"Unsupported method")
	}
}


func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/")]
	isbn = isbn   // A hack

	urlString := r.URL.String()
	isbnString := strings.TrimLeft(urlString, "/api/books/")
	log.Printf("urlString = %v , ISBN = %s", urlString, isbnString)

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbnString)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		log.Println("case http.MethodPut")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("http.StatusInternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJson(body)
		log.Printf(" Book : %v", book)
		exists := UpdateBook(isbnString, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbnString)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	log.Printf(" UpdateBook : %v", book)
	_, exists := books[isbn]
	if exists {
		log.Printf(" it exists : %v\n", book)
		books[isbn] = book
		log.Print(" Updating books slice", book)
	}
	return exists
}


// GetBook returns the  book for a given isbn
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}


func CreateBook(book Book) (string, bool) {
	books[book.ISBN] = book
	return book.ISBN, true
}

// Book version
func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}


/*
// My Version
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type","application/json; charset-utf-8")
	w.Write(b)
}

*/


var Books = []Book{
	Book{Title:"The Hitchhiker's Guide to the Galaxy", Author:"Douglas Adams",ISBN:"0-671-62964-6"},
	Book{Title:"Cloud Native Go", Author:"M.L.Reimer",ISBN:"0123456"},
	Book{Title:"A Spell for Chameleon", Author:"Piers Anthony",ISBN:"0-345-25855-X"},
}


func (b Book) toJSON() []byte {
	TOJSON, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return TOJSON
}

func FromJson(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		log.Fatal(err)
	}
	return book
}



func BooksHandleFuncOriginal(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type","application/json; charset-utf-8")
	w.Write(b)
}


