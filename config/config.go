package config

import (
	"database/sql"
	"net/http"
	book2 "tidy/internal/adapters/api/book"
	"tidy/internal/adapters/db/book"
	book1 "tidy/internal/domain/book"
)

func Config(db *sql.DB) *http.ServeMux {

	router := http.NewServeMux()
	bookStorage := book.NewStorage(db)

	bookService := book1.NewService(bookStorage)
	bookHandler := book2.NewHandler(bookService)

	bookHandler.Register(router)

	return router
}
