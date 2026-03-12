package handler

import (
	"net/http"
	"encoding/json"
	"go-book-tracker/model"
	"go-book-tracker/service"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	books := h.service.GetBooks()
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request){
	var newBook model.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.AddBook(newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	book, found := h.service.GetBookByID(id)

	if !found{
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	var updatedBook model.Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook.ID = id

	success := h.service.UpdateBook(id, updatedBook)
    
	if !success {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request){

	id := chi.URLParam(r, "id")

	success := h.service.DeleteBook(id)

	if !success {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

