package handler

import (
	"net/http"
	"encoding/json"
	"go-book-tracker/model"
	"go-book-tracker/storage"
	"github.com/go-chi/chi/v5"
)

func BookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		GetBooks(w, r)
		return
	}

	if r.Method == http.MethodPost {
		AddBook(w, r)
		return
	}

	http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
}

func BookHandlerByID(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		GetBookByID(w, r)
		return
	}
    if r.Method == http.MethodPut{
		UpdateBook(w, r)
		return
	}
	if r.Method == http.MethodDelete{
		DeleteBook(w, r)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(storage.Books)
}

func AddBook(w http.ResponseWriter, r *http.Request){
	var newBook model.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.Books = append(storage.Books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	for _, book := range storage.Books {
		if book.ID == id{
			json.NewEncoder(w).Encode(book)
			return
		}
	}	

	http.Error(w, "Book Not Found", http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	var updatedBook model.Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i,book := range storage.Books {
		if book.ID == id {

            updatedBook.ID = id
			storage.Books[i] = updatedBook

			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book Not Found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){

	id := chi.URLParam(r, "id")

	for i,book := range storage.Books {
		if book.ID == id {
            storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book Not Found", http.StatusNotFound)
}

