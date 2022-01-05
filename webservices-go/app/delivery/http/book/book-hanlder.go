package httpHandler

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	commonHTTP "webserver/app/delivery/http"
	"webserver/app/entity"
)

type BookHanlder struct {
	book Repository
}

func NewArticleHttpHandler(repo Repository) *BookHanlder {
	return &BookHanlder{
		book: repo,
	}
}

func (hanlder *BookHanlder) HandleFunc() {
	http.HandleFunc("/v1/books", hanlder.Books) // update this line of code.
	// http.HandleFunc("/v1/books/{id}", hanlder.DeleteBook)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Starting server at port 8080\n")
}

func (hanlder *BookHanlder) Books(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method, r)
	switch r.Method {
	case "GET":
		hanlder.SearchBooks(w, r)
	case "POST":
		hanlder.CreateBooks(w, r)
	case "DELETE":
		hanlder.DeleteBook(w, r)
	default:
		fmt.Print(r.Method)
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (hanlder *BookHanlder) GetBook(id entity.ID) (*entity.Book, error) {
	return &entity.Book{}, nil
}

func (hanlder *BookHanlder) SearchBooks(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Search")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Encoding", "gzip")
	// time.Sleep(time.Second * 4)
	if r.Method != "GET" {
		http.Error(w, "Method is not support!", http.StatusNotFound)
		return
	}
	keys := r.URL.Query()["keyword"] // get keyword param
	var keyword string = ""
	if len(keys) > 0 {
		keyword = keys[0]
	}
	data, err := hanlder.book.SearchBooks(keyword)
	if err != nil {
		// not found
		var data []*entity.Book
		w.Write(gzipText(&RResponse{Code: 404, Data: data}))
		return
	}
	w.Write(gzipText(&RResponse{Code: 200, Data: data}))
}

func (hanlder *BookHanlder) ListBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	return books, nil
}

func (hanlder *BookHanlder) UpdateBook(e *entity.Book) error {
	return nil
}

func (hanlder *BookHanlder) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	keys := r.URL.Query() // get keyword param
	fmt.Print("keys: ", keys)
	js, _ := json.Marshal("")
	w.Write(js)
}

func (hanlder *BookHanlder) CreateBooks(w http.ResponseWriter, r *http.Request) {
	if commonHTTP.HasContentType(r, "application/json") {
		var book RBook
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = hanlder.book.CreateBook(book.Title, book.Author, book.Pages, book.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Apply multipart
		r.ParseMultipartForm(0)
		var m []RBook
		if err := json.Unmarshal([]byte(r.FormValue("book")), &m); err != nil {
			panic(err)
		}
		for i := 0; i < len(m); i++ {
			hanlder.book.CreateBook(m[i].Title, m[i].Author, m[i].Pages, m[i].Quantity)
		}

	}

	js, err := json.Marshal(RResponse{Code: 200, Data: make([]*entity.Book, 0)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func gzipText(data *RResponse) []byte {
	box := data
	b, _ := json.Marshal(box)

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	defer zw.Close()
	if _, err := zw.Write(b); err != nil {
		panic(err)
	}
	zw.Close()
	return buf.Bytes()
}

type RResponse struct {
	Code int            `json:"code"`
	Data []*entity.Book `json:"data"`
}

type RBook struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Pages    int    `json:"pages"`
	Quantity int    `json:"quantity"`
}
