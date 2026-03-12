package service

import (
	"go-book-tracker/model"
	"go-book-tracker/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks() []model.Book {
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id string) (*model.Book, bool) {
	return s.repo.GetByID(id)
}

func (s *BookService) AddBook(newBook model.Book) {
	s.repo.Create(newBook)
}

func (s *BookService) UpdateBook(id string, updatedBook model.Book) bool {
	return s.repo.Update(id, updatedBook)
}

func (s *BookService) DeleteBook(id string) bool {
	return s.repo.Delete(id)
}