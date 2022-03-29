package dbase

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CheckDB() *sql.DB {
	_, err := os.Stat("dbase/database-sqlite.db")
	if os.IsNotExist(err) {
		createFile()
	}
	var d Database
	d.open("dbase/database-sqlite.db")
	d.createTable()
	return d.db
}

func createFile() {
	file, err := os.Create("dbase/database-sqlite.db")
	if err != nil {
		log.Fatalf("file doesn't create %v", err)
	}
	defer file.Close()
}

func (d *Database) open(file string) {
	var err error
	d.db, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("this error is in dbase/open() %v", err)
	}
}

func (d *Database) createTable() {
	_, err := d.db.Exec(`CREATE TABLE IF NOT EXISTS book (
        "id"    INTEGER NOT NULL UNIQUE,
        "title"    TEXT NOT NULL UNIQUE,
        "description"    TEXT NOT NULL,
        "image"    TEXT,
		"author" TEXT NOT NULL,
		"ganre"	TEXT NOT NULL,
        PRIMARY KEY("id" AUTOINCREMENT)
    );`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE users")
	}

}
