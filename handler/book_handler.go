package handler

import (
	"net/http"
	"encoding/json"
	"go-book-tracker/model"
	"go-book-tracker/service"
	"go-book-tracker/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	books := h.service.GetBooks(ctx)
	utils.JSON(w, http.StatusOK, books)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()

	var newBook model.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.ValidateBook(newBook)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	newBook.ID = uuid.New().String()
	h.service.AddBook(ctx, newBook)
	utils.JSON(w, http.StatusCreated, newBook)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	id := chi.URLParam(r, "id")

	book, found := h.service.GetBookByID(ctx, id)

	if !found{
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.JSON(w, http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var updatedBook model.Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.ValidateBook(updatedBook)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedBook.ID = id

	success := h.service.UpdateBook(ctx, id, updatedBook)
    
	if !success {
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.JSON(w, http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()

	id := chi.URLParam(r, "id")

	success := h.service.DeleteBook(ctx, id)

	if !success {
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)

}

