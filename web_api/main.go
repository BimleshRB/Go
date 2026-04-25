package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 1. DATA MODEL
// Use 'json' tags to control how fields are serialized.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// In-memory data store
var (
	books = []Book{
		{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
		{ID: "2", Title: "Clean Code", Author: "Robert C. Martin"},
	}
	mu sync.Mutex // To handle concurrent access (from our concurrency lesson!)
)

// 2. HANDLERS
// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Create a new book (POST)
func createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	books = append(books, newBook)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// 3. MIDDLEWARE
// A simple logger middleware
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	fmt.Println("--- Web Development in Go ---")
	fmt.Println("Starting server on :8080...")

	// 4. ROUTING
	mux := http.NewServeMux()

	// Registering routes with middleware
	mux.HandleFunc("/books", logger(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w, r)
		case http.MethodPost:
			createBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Go Book Store API!")
		fmt.Fprintln(w, "Available endpoints:")
		fmt.Fprintln(w, "- GET /books")
		fmt.Fprintln(w, "- POST /books")
	})

	// 5. START SERVER
	// Using log.Fatal to catch errors and exit
	log.Fatal(http.ListenAndServe(":8080", mux))

	/*
	SUMMARY:
	- 'net/http' is the heart of web development in Go.
	- 'json' tags allow for easy JSON mapping.
	- 'http.ServeMux' is a simple but powerful router.
	- Handlers are functions with (w http.ResponseWriter, r *http.Request).
	- Middleware can be used to wrap handlers for shared logic (logging, auth).
	- Go servers are concurrent by default! Every request runs in its own goroutine.
	*/
}
