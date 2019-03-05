package main

import (
	"fmt"
	"getstrtCloudNative/api"
	"log"
	"net/http"
	"os"
)

func main() {


	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/bookorginal", api.BooksHandleFuncOriginal)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	//
	//http.HandleFunc("/api/books", api.BooksHandleFunc)
	//http.HandleFunc("/api/books/", api.BookHandleFunc)



	http.ListenAndServe(port(),nil)

}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Context-type","text/plain")
	fmt.Fprintf(w,message)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Println("Port = :"+port)
	return ":"+port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Hello Cloud Native Go.")

}




