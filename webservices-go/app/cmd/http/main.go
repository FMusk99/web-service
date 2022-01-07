package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"webserver/app/config"

	_postgresRepo "webserver/app/infrastructure/repository/postgres"
	_bookUsecase "webserver/app/usecase/book"
	_categoryUsecase "webserver/app/usecase/category"

	_bookHTTPDelivery "webserver/app/delivery/http/book"
	_categoryHTTPDelivery "webserver/app/delivery/http/category"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func checkError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func initualRoute(db *sql.DB) {
	bookRepo := _postgresRepo.NewBookSQL(db)
	// bookRepo.InitialDB()
	bookService := _bookUsecase.NewService(bookRepo)
	webService := _bookHTTPDelivery.NewArticleHttpHandler(bookService)
	webService.HandleFunc()

	categoryRepo := _postgresRepo.NewCategorySQL(db)
	// categoryRepo.InitialDB()
	categoryService := _categoryUsecase.NewService(categoryRepo)
	webService1 := _categoryHTTPDelivery.NewCategoryHttpHandler(categoryService)
	webService1.HandleFunc()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	checkError(err)
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer db.Close()
	initualRoute(db)
}
