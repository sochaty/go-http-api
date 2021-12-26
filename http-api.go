package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


// Post //
type Book struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Author User `json:"author"`
}

// User //
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email string `json:"email"`
}

var books []Book = []Book{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", patchBook).Methods("PATCH")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":5000", router)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}
	book := books[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	// routeVariable := mux.Vars(r)["item"]
	books = append(books, newBook)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	books[id] = updatedBook
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func patchBook(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}
	book := &books[id]
	json.NewDecoder(r.Body).Decode(book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}

	books = append(books[:id], books[id+1:]...)
	w.WriteHeader(200)
}