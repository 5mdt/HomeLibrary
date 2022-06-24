package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the LibraryLog! ")
	fmt.Fprintf(w, "It is a logger of places and borrowers of books in your library")
	fmt.Println("Endpoint Hit: Main")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllBooks)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("LibraryLog Rest API")
	Books = []Book{
		Book{Title: "Alice in Wonderland", Author: "Lewis Carrol", Desc: "Book Description", Content: "Book Content"},
		Book{Title: "Hello 2", Desc: "Book Description", Content: "Book Content"},
	}
	handleRequests()
}
