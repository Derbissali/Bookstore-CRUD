package book

import (
	"database/sql"
	"fmt"
	"log"
	"tidy/internal/domain/book"
	"tidy/internal/model"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) book.BookStorage {
	return &bookStorage{
		db: db,
	}
}

func (c *bookStorage) Create(m model.Book) error {
	_, err := c.db.Exec(`INSERT INTO book (title, description, image, author, ganre) VALUES (?, ?, ?, ?, ?)`, m.Title, m.Description, m.Image, m.Author, m.Ganre)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *bookStorage) Read() ([]model.Book, error) {
	rows, err := c.db.Query(`SELECT book.id,book.title,book.description,book.ganre, book.author, book.Image
	FROM book`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var m model.Book
	for rows.Next() {

		var a model.Book
		err = rows.Scan(&a.ID, &a.Title, &a.Description, &a.Ganre, &a.Author, &a.Image)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func (c *bookStorage) ReadOne(id string) ([]model.Book, error) {
	rows, err := c.db.Query(`SELECT book.id,book.title,book.description,book.ganre, book.author, book.Image
	FROM book WHERE book.id=?`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var m model.Book
	for rows.Next() {

		var a model.Book
		err = rows.Scan(&a.ID, &a.Title, &a.Description, &a.Ganre, &a.Author, &a.Image)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}

func (c *bookStorage) Update(m model.Book) error {
	_, err := c.db.Exec(`UPDATE book SET title=?, description=?, image=?, author=?, ganre=? WHERE id = ?`, m.Title, m.Description, m.Image, m.Author, m.Ganre, m.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *bookStorage) Delete(id int) error {
	_, err := c.db.Exec(`DELETE FROM book WHERE book.id=?`, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
