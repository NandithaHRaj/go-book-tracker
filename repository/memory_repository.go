package repository

import "go-book-tracker/model"

type BookRepository interface {
	GetAll() []model.Book
	GetByID(string) (*model.Book, bool)
	Create(model.Book) 
	Update(string, model.Book) bool
	Delete(string) bool
}

type MemoryRepository struct {
	books []model.Book
}

func (m *MemoryRepository) GetAll() []model.Book {
	return m.books
}

func (m *MemoryRepository) GetByID(id string) (*model.Book, bool)			 {
	for _,book := range m.books {
		if book.ID == id{
			return &book, true
		}
	}
    return nil, false
}

func (m *MemoryRepository) Create(book model.Book) {
	m.books = append(m.books, book)
}

func (m *MemoryRepository) Update(id string, updatedBook model.Book) bool {
    for i,book := range m.books {
		if book.ID == id{
			m.books[i] = updatedBook
			return true
		}
	}
	return false
}

func (m *MemoryRepository) Delete(id string) bool {
	 for i,book := range m.books {
		if book.ID == id{
			m.books = append(m.books[:i], m.books[i+1:]...)
			return true
		}
	}
	return false 
}