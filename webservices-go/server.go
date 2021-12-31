package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"webserver/db"

	"bytes"
	"compress/gzip"
)

func main() {

	//connecting db
	db.ConnectDb()
	db.GetUser("jo")
	http.HandleFunc("/hello", helloHandler) // update this line of code.
	//http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/search", searchByKeywordHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// func GetTest() {
// 	panic("unimplemented")
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not support!", http.StatusNotFound)
		return
	}

	//fmt.Fprintf(w, "Hello!")
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 3)
	if r.URL.Path != "/search" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not support!", http.StatusNotFound)
		return
	}

	c := []Candidate{
		{Id: 1, Name: "Tom", Interests: []string{"art", "coding", "music", "travel"}, Language: "golang", Experience: false},
		{Id: 2, Name: "Jerry", Interests: []string{"art", "coding", "music", "travel"}, Language: "golang", Experience: false},
		{Id: 3, Name: "Z", Interests: []string{"art", "coding", "music", "travel"}, Language: "golang", Experience: true},
		{Id: 3, Name: "", Interests: []string{"art", "coding", "music", "travel"}, Language: "golang", Experience: true},
	}
	founded := []Candidate{}         // add new array of result data
	keys := r.URL.Query()["keyword"] // get keyword param
	// fmt.Println("Request: ", r.URL, keys, keys[0])
	for _, v := range c {
		if v.Name == keys[0] { // is Name property equal keywork
			founded = append(founded, v)
			// fmt.Println("Found: ", v.Name )
			break
		}
	}

	var resposeData ResponseDataPattern
	if len(founded) > 0 {
		resposeData = ResponseDataPattern{
			Code: 200,
			Data: founded,
		}

	} else {
		resposeData = ResponseDataPattern{
			Code: 404,
			Data: []Candidate{},
		}
	}
	js, err := json.Marshal(resposeData)
	//js, err := json.Marshal(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func searchByKeywordHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not support!", http.StatusNotFound)
		return
	}
	keys := r.URL.Query()["keyword"] // get keyword param
	founded := db.GetUser(keys[0])
	var resposeData UserData
	if len(founded) > 0 {
		resposeData = UserData{
			Code: 200,
			Data: founded,
		}

	} else {
		resposeData = UserData{
			Code: 404,
			Data: []db.UserModel{},
		}
	}
	// js, err := json.Marshal(resposeData)
	// fmt.Println(js)
	//js, err := json.Marshal(c)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Encoding", "gzip")
	time.Sleep(time.Second * 4)
	w.Write(gzipText(resposeData))
	// w.Write(js)
}

func gzipText(data UserData) []byte {
	// var buf bytes.Buffer
	// zw := gzip.NewWriter(&buf)

	// // Setting the Header fields is optional.
	// zw.Name = "a-new-hope.txt"
	// zw.Comment = "an epic space opera by George Lucas"
	// zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	// _, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := zw.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// zr, err := gzip.NewReader(&buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := io.Copy(os.Stdout, zr); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := zr.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	box := data
	// box.ars = make([]int, 100)
	// for i := range box.ars {
	//     box.ars[i] = i
	// }
	b, _ := json.Marshal(box)

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	defer zw.Close()
	if _, err := zw.Write(b); err != nil {
		panic(err)
	}
	zw.Close()
	fmt.Println(zw)
	return buf.Bytes()
}

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}
type Byte1 struct {
	Byte []int
}

type Profile struct {
	Name    string
	Hobbies []string
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getUsers() (res AllUsers, errString string) {
	file, err := os.OpenFile("list.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var alUsrs AllUsers
	json.Unmarshal(b, &alUsrs.Users)
	checkError(err)
	return alUsrs, ""
}

type User struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Balance   float64 `json:"balance"`
}

type AllUsers struct {
	Users []*User
}

type Candidate struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Interests  []string `json:"interests"`
	Language   string   `json:"language"`
	Experience bool     `json:"experience"`
}

type Message struct {
	Name string
	Body string
	Time int64
	Foo  string
}

type ResponseDataPattern struct {
	Code int         `json:"code"`
	Data []Candidate `json:"data"`
}

type UserData struct {
	Code int            `json:"code"`
	Data []db.UserModel `json:"data"`
}
