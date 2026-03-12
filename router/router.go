package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"go-book-tracker/handler"
)

func SetupRouter(handler *handler.BookHandler) http.Handler {

	r := chi.NewRouter()

	r.Get("/books", handler.GetBooks)
	r.Post("/books", handler.AddBook)
	r.Get("/books/{id}", handler.GetBookByID)
	r.Put("/books/{id}", handler.UpdateBook)
	r.Delete("/books/{id}", handler.DeleteBook)

	return r
}