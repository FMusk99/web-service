package http

import (
	"database/sql"
	"fmt"
	"mime"
	"net/http"
	"regexp"
	"strings"
	// bookHanlder "webserver/app/delivery/http/book"
	// categoryHanlder "webserver/app/delivery/http/category"
)

type HttpServer struct {
	db *sql.DB
}

func NewRestServer(db *sql.DB) *HttpServer {
	return &HttpServer{
		db: db,
	}
}

var categories = regexp.MustCompile(`^/v1/categories*`) //categories resource
var books = regexp.MustCompile(`^/v1/books*`)           //books resource

// Determine whether the request `content-type` includes a
// server-acceptable mime-type

// Failure should yield an HTTP 415 (`http.StatusUnsupportedMediaType`)
func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		fmt.Println(v, mimetype)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

func Router(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL.Path)
	switch {
	case categories.MatchString(r.URL.Path):
		w, r = categoriesResouce(w, r)
		fmt.Print(w, r)
	case books.MatchString(r.URL.Path):
		_, _ = booksResouce(w, r)
	default:
		w.Write([]byte("404 page not found."))
	}
}

func categoriesResouce(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	return w, r
}

func booksResouce(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	return w, r
}

func (h *HttpServer) InitialRouter() {
	// 	bookRepo := repository.NewBookSQL(h.db)
	// 	bookRepo.InitialDB()
	// 	bookService := book.NewService(bookRepo)
	// 	webService := bookHanlder.NewArticleHttpHandler(bookService)
	// 	// webService.HandleFunc()

	// 	categoryRepo := repository.NewCategorySQL(h.db)
	// 	categoryRepo.InitialDB()
	// 	categoryService := category.NewService(categoryRepo)
	// 	webService1 := categoryHanlder.NewCategoryHttpHandler(categoryService)
	// 	// webService1.HandleFunc()

	// 	if err := http.ListenAndServe(":8080", nil); err != nil {
	// 		log.Fatal(err)
	// 	}
}
