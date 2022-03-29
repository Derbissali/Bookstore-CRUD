package book

import (
	"fmt"
	"tidy/internal/adapters/api"
	"tidy/internal/model"
)

type service struct {
	storage BookStorage
}

func NewService(storage BookStorage) api.BookService {
	return &service{
		storage: storage,
	}
}

func (s *service) Create(m model.Book) error {
	err := s.storage.Create(m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *service) Read() ([]model.Book, error) {
	m, err := s.storage.Read()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return m, nil
}
func (s *service) ReadOne(id string) ([]model.Book, error) {
	m, err := s.storage.ReadOne(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return m, nil
}
func (s *service) Update(m model.Book) error {

	err := s.storage.Update(m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *service) Delete(m model.Book) error {
	err := s.storage.Delete(m.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
