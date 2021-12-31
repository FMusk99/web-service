package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	commonHTTP "webserver/app/delivery/http"
	"webserver/app/entity"
)

type CategoryHanlder struct {
	category Repository
}

func NewCategoryHttpHandler(repo Repository) *CategoryHanlder {
	return &CategoryHanlder{
		category: repo,
	}
}

func (hanlder *CategoryHanlder) HandleFunc() {
	http.HandleFunc("/v1/categories", hanlder.Categories) // update this line of code.
}

func (hanlder *CategoryHanlder) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		hanlder.SearchCategories(w, r)
	case "POST":
		hanlder.CreateCategories(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (hanlder *CategoryHanlder) GetBook(id entity.ID) (*entity.Category, error) {
	return &entity.Category{}, nil
}

func (hanlder *CategoryHanlder) SearchCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	data, err := hanlder.category.SearchCategories(keyword)
	if err != nil {
		// not found
		var data []*entity.Category
		js, err := json.Marshal(RResponse{Code: 404, Data: data})
		if err != nil {
			panic(err)
		}
		w.Write(js)
		return
	}
	js, err := json.Marshal(RResponse{Code: 200, Data: data})
	if err != nil {
		panic(err)
	}
	w.Write(js)
}

func (hanlder *CategoryHanlder) ListBooks() ([]*entity.Book, error) {
	var books []*entity.Book
	return books, nil
}

func (hanlder *CategoryHanlder) UpdateBook(e *entity.Book) error {
	return nil
}

func (hanlder *CategoryHanlder) DeleteBook(id entity.ID) error {
	return nil
}

func (hanlder *CategoryHanlder) CreateCategories(w http.ResponseWriter, r *http.Request) {

	if commonHTTP.HasContentType(r, "application/json") {
		var category RDescription
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = hanlder.category.CreateCategory(category.Name, category.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else { // Apply multipart
		var categories []RDescription
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&categories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Do something with the Person struct...
		r.ParseMultipartForm(0)
		for i := 0; i < len(categories); i++ {
			_, err := hanlder.category.CreateCategory(categories[i].Name, categories[i].Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
	}

	js, err := json.Marshal(RResponse{Code: 200, Data: make([]*entity.Category, 0)})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Done")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

type RResponse struct {
	Code int                `json:"code"`
	Data []*entity.Category `json:"data"`
}

type RDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
