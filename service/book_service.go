package service

import (
	"go-book-tracker/model"
	"go-book-tracker/repository"
	"context"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks(ctx context.Context) []model.Book {
	return s.repo.GetAll(ctx)
}

func (s *BookService) GetBookByID(ctx context.Context, id string) (*model.Book, bool) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) AddBook(ctx context.Context, newBook model.Book) {
	s.repo.Create(ctx, newBook)
}

func (s *BookService) UpdateBook(ctx context.Context, id string, updatedBook model.Book) bool {
	return s.repo.Update(ctx, id, updatedBook)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) bool {
	return s.repo.Delete(ctx, id)
}