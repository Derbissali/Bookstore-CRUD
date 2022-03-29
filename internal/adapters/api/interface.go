package api

import (
	"net/http"
	"tidy/internal/model"
)

type Handler interface {
	Register(router *http.ServeMux)
}
type BookService interface {
	Create(m model.Book) error
	Read() ([]model.Book, error)
	ReadOne(id string) ([]model.Book, error)
	Update(m model.Book) error
	Delete(m model.Book) error
}
