package book

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"tidy/internal/adapters/api"
	"tidy/internal/model"
	"time"
)

type handlerBook struct {
	bookService api.BookService
}

func NewHandler(service api.BookService) api.Handler {
	return &handlerBook{
		bookService: service,
	}
}
func (h *handlerBook) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.home_page)
	router.HandleFunc("/addBook", h.addBook)
	router.HandleFunc("/book/", h.Bookpage)
	router.HandleFunc("/update", h.UpdateBook)
	router.HandleFunc("/delete", h.DeleteBook)

}

func (h *handlerBook) home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")

	if err != nil {
		log.Printf("Error main-page html Post Handler GetAll method:--> %v\n", err)
		return
	}
	var M model.Book
	M.Rows, err = h.bookService.Read()
	if err != nil {
		log.Printf("ERROR book handler BookRead method:--> %v\n", err)

		return
	}
	err = temp.Execute(w, M)

	if err != nil {
		log.Printf("ERROR post handler GetAll method Execute:---> %v\n", err)

		return
	}
}

func (h *handlerBook) addBook(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/addpost.html")
	if err != nil {
		log.Printf("Error main-page html Post Handler GetAll method:--> %v\n", err)
		return
	}
	var M model.Book

	switch r.Method {
	case "GET":
		if err != nil {
			log.Printf("ERROR post handler PostCreate method SelectCategory function:--> %v\n", err)

			return
		}
		temp.Execute(w, M)
	case "POST":
		r.ParseMultipartForm(0)

		file, handler, err := r.FormFile("myFile")
		var fileName string
		if err == nil {
			dst, err := os.Create(fmt.Sprintf("assets/temp-images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer dst.Close()
			fileName = strings.TrimPrefix(dst.Name(), "assets/temp-images/")
			if len(fileName) == 0 {
				fileName = "1"
			}
			fmt.Println(fileName) // Copy the uploaded file to the filesystem
			// at the specified destination
			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//func() {}
		}
		argument := model.Book{
			Title:       r.FormValue("titleB"),
			Description: r.FormValue("descriptionB"),
			Author:      r.FormValue("authorB"),
			Ganre:       r.FormValue("genreB"),
			Image:       fileName,
		}

		h.bookService.Create(argument)

		http.Redirect(w, r, "/", 301)
		return
	}
}
func (h *handlerBook) Bookpage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/book_page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

	id := r.RequestURI[6:]
	var M model.Book
	M.Rows, err = h.bookService.ReadOne(id)
	tmpl.Execute(w, M)
}

func (h *handlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("./templates/post_page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.ParseMultipartForm(0)

	file, handler, err := r.FormFile("myFile")
	var fileName string
	if err == nil {
		dst, err := os.Create(fmt.Sprintf("assets/temp-images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()
		fileName = strings.TrimPrefix(dst.Name(), "assets/temp-images/")
		if len(fileName) == 0 {
			fileName = "1"
		}
		fmt.Println(fileName) // Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//func() {}
	}
	id, _ := strconv.Atoi(r.FormValue("idwka"))
	argument := model.Book{
		ID:          id,
		Title:       r.FormValue("titleB"),
		Description: r.FormValue("descriptionB"),
		Author:      r.FormValue("authorB"),
		Ganre:       r.FormValue("genreB"),
		Image:       fileName,
	}
	err = h.bookService.Update(argument)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/", 301)
	return
}
func (h *handlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("./templates/post_page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("idwka"))
	argument := model.Book{
		ID: id,
	}
	err = h.bookService.Delete(argument)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/", 301)
	return
}
