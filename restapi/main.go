package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book

func getBooks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(books)
}

func getBook(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}

	json.NewEncoder(resp).Encode(&Book{})
}

func createBook(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	json.NewEncoder(resp).Encode(books)
}

func updateBook(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}

	json.NewEncoder(resp).Encode(&Book{})
}

func deleteBook(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}

	json.NewEncoder(resp).Encode(&Book{})
}

func main() {
	//Init Router
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Risalah Cinta",
		Author: &Author{FirstName: "Riyansyah", LastName: "Iqbal"}})

	books = append(books, Book{ID: "2", Isbn: "12346", Title: "Risalah Hati",
		Author: &Author{FirstName: "Riyansyah", LastName: "Iqbal"}})

	books = append(books, Book{ID: "3", Isbn: "12347", Title: "Risalah Jiwa",
		Author: &Author{FirstName: "Riyansyah", LastName: "Iqbal"}})

	router.HandleFunc("/api/books", getBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", getBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books", createBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", deleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
