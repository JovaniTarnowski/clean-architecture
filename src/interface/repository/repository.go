package repository

import "github.com/jovanitarnowski/clean-architecture/src/domain"

type DBHandler interface {
	FindAllBooks() ([]*domain.Book, error)
	SaveBook(book domain.Book) error
	SaveAuthor(author domain.Author) error
}
