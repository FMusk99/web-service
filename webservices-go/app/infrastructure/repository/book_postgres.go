package repository

import (
	"database/sql"
	"fmt"
	"time"
	"webserver/app/entity"
)

//BookSQL sql repo
type BookSQL struct {
	db *sql.DB
}

//NewBookSQL create new repository
func NewBookSQL(db *sql.DB) *BookSQL {
	return &BookSQL{
		db: db,
	}
}

//Create a book
func (r *BookSQL) Create(e *entity.Book) (entity.ID, error) {
	sql_statement := "INSERT INTO book (id, title, author, pages, quantity, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	_, err := r.db.Exec(sql_statement, e.ID, e.Title, e.Author, e.Pages, e.Quantity, time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02"))
	if err != nil {
		return e.ID, err
	}
	// err = r.db.Close()
	return e.ID, nil
}

//Get a book
func (r *BookSQL) Get(id entity.ID) (*entity.Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Book
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
	}
	return &b, nil
}

//Update a book
func (r *BookSQL) Update(e *entity.Book) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = ?", e.Title, e.Author, e.Pages, e.Quantity, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search books
func (r *BookSQL) Search(query string) ([]*entity.Book, error) {
	sql_statement := "SELECT * FROM book where title ~* '" + query + "';"
	rows, err := r.db.Query(sql_statement)
	var book []*entity.Book
	if err != nil {
		fmt.Print("Query error")
		panic(err)
	}
	// defer rows.Close()
	var id entity.ID
	var title string
	var author string
	var pages int
	var quantity int
	var updated_at string
	var created_at string
	for rows.Next() {
		switch err := rows.Scan(&id, &title, &author, &pages, &quantity, &created_at, &updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			book = append(book, &entity.Book{ID: id, Title: title, Author: author, Pages: pages, Quantity: quantity, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		default:
			if err != nil {
				return book, err
			}
		}
	}
	return book, nil
}

//List books
func (r *BookSQL) List() ([]*entity.Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var books []*entity.Book
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

//Delete a book
func (r *BookSQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from book where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (w *BookSQL) InitialDB() {
	// Drop previous table of same name if one exists.
	_, err := w.db.Exec("DROP TABLE IF EXISTS book;")
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished dropping table (if existed)")
	// Create table.
	_, err = w.db.Exec("CREATE TABLE book (id VARCHAR(50) PRIMARY KEY, title VARCHAR(50), author VARCHAR(100), Pages INTEGER,Quantity INTEGER, updated_at VARCHAR(200), created_at VARCHAR(200));")
	if err != nil {
		panic(err)
	}
	fmt.Print("Finished creating table: Book \n")
}

type IBookRepository interface {
	Create(e *entity.Book) (entity.ID, error)
	Get(id entity.ID) (*entity.Book, error)
	Update(e *entity.Book) error
	Search(query string) ([]*entity.Book, error)
	List() ([]*entity.Book, error)
	Delete(id entity.ID) error
	InitialDB()
}
