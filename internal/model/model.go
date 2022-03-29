package model

type Book struct {
	ID          int
	Title       string
	Description string
	Author      string
	Ganre       string
	Image       string
	Rows        []Book
}
