package handler

import (
	"net/http"
	"encoding/json"
	"go-book-tracker/model"
	"go-book-tracker/service"
	"go-book-tracker/utils"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request){
	books := h.service.GetBooks()
	utils.JSON(w, http.StatusOK, books)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request){
	var newBook model.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.AddBook(newBook)
	utils.JSON(w, http.StatusCreated, newBook)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	book, found := h.service.GetBookByID(id)

	if !found{
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	utils.JSON(w, http.StatusOK, book)
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

	utils.JSON(w, http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request){

	id := chi.URLParam(r, "id")

	success := h.service.DeleteBook(id)

	if !success {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)

}

