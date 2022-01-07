package http

import (
	"net/http"
	"webserver/app/entity"
)

type Reader interface {
	GetBook(id entity.ID) (*entity.Book, error)
	SearchBooks(query string) ([]*entity.Book, error)
	ListBooks() ([]*entity.Book, error)
}

//Writer book writer
type Writer interface {
	CreateBook(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(id string) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type IbookHandler interface {
	GetBook(id entity.ID) (*entity.Book, error)
	SearchBooks(w http.ResponseWriter, r *http.Request)
	ListBooks() ([]*entity.Book, error)
	CreateBook(title string, author string, pages int, quantity int) (entity.ID, error)
	CreateBooks() (string, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
