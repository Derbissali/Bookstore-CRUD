package main

import (
	"log"
	"net/http"
	"tidy/config"
	"tidy/dbase"
)

func main() {
	db := dbase.CheckDB()

	router := config.Config(db)

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("port: 8080 is listening")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("%v", err)
		// log.Fatal("ListenAndServe ERROR")
	}
}
