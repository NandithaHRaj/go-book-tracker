package repository

import (
	"go-book-tracker/model"
	"context"
)

type BookRepository interface {
	GetAll(context.Context) []model.Book
	GetByID(context.Context, string) (*model.Book, bool)
	Create(context.Context, model.Book) 
	Update(context.Context, string, model.Book) bool
	Delete(context.Context, string) bool
}

type MemoryRepository struct {
	books []model.Book
}

func (m *MemoryRepository) GetAll(ctx context.Context) []model.Book {
	return m.books
}

func (m *MemoryRepository) GetByID(ctx context.Context, id string) (*model.Book, bool)			 {
	for i := range m.books {
		if m.books[i].ID == id{
			return &m.books[i], true
		}
	}
    return nil, false
}

func (m *MemoryRepository) Create(ctx context.Context, book model.Book) {
	m.books = append(m.books, book)
}

func (m *MemoryRepository) Update(ctx context.Context, id string, updatedBook model.Book) bool {
    for i,book := range m.books {
		if book.ID == id{
			m.books[i] = updatedBook
			return true
		}
	}
	return false
}

func (m *MemoryRepository) Delete(ctx context.Context, id string) bool {
	 for i,book := range m.books {
		if book.ID == id{
			m.books = append(m.books[:i], m.books[i+1:]...)
			return true
		}
	}
	return false 
}