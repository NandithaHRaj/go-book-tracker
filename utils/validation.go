package utils

import (
	"errors"
	"strings"
	"go-book-tracker/model"
)

func ValidateBook(book model.Book) error {

	if strings.TrimSpace(book.Title) == "" {
		return errors.New("Title is required")
	}

	if strings.TrimSpace(book.Author) == "" {
		return errors.New("Author is required")
	}

	if book.Rating < 0 || book.Rating > 5 {
		return errors.New("Rating must be between 0 and 5")
	}

	return nil

}
