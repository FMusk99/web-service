package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"webserver/app/config"
	bookHanlder "webserver/app/delivery/http/book"
	categoryHanlder "webserver/app/delivery/http/category"
	"webserver/app/infrastructure/repository"
	"webserver/app/usecase/book"
	"webserver/app/usecase/category"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func checkError(err error) {
	if err != nil {
		// panic(err)
		fmt.Print(err)
	}
}

func main() {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)
	db, err := sql.Open("postgres", dataSourceName)
	fmt.Print(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer db.Close()
	bookRepo := repository.NewBookSQL(db)
	bookRepo.InitialDB()
	bookService := book.NewService(bookRepo)
	webService := bookHanlder.NewArticleHttpHandler(bookService)
	webService.HandleFunc()

	categoryRepo := repository.NewCategorySQL(db)
	categoryRepo.InitialDB()
	categoryService := category.NewService(categoryRepo)
	webService1 := categoryHanlder.NewCategoryHttpHandler(categoryService)
	webService1.HandleFunc()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
