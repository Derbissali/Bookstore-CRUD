package book

import "tidy/internal/model"

type BookStorage interface {
	Create(m model.Book) error
	Read() ([]model.Book, error)
	ReadOne(id string) ([]model.Book, error)
	Update(m model.Book) error
	Delete(id int) error
}
