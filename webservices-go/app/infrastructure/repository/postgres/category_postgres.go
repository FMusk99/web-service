package postgres

import (
	"database/sql"
	"fmt"
	"time"
	"webserver/app/entity"
)

//CategorySQL sql repo
type CategorySQL struct {
	db *sql.DB
}

//NewCategorySQL create new repository
func NewCategorySQL(db *sql.DB) *CategorySQL {
	return &CategorySQL{
		db: db,
	}
}

//Create a book
func (r *CategorySQL) Create(e *entity.Category) (entity.ID, error) {
	sql_statement := "INSERT INTO category (id, name, updated_at, created_at) VALUES ($1, $2, $3, $4);"
	_, err := r.db.Exec(sql_statement, e.ID, e.Name, time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02"))
	if err != nil {
		return e.ID, err
	}
	// err = r.db.Close()
	return e.ID, nil
}

//Get a book
func (r *CategorySQL) Get(id entity.ID) (*entity.Category, error) {
	stmt, err := r.db.Prepare(`select id, name, created_at from book where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Category
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Name, &b.CreatedAt)
	}
	return &b, nil
}

//Update a book
func (r *CategorySQL) Update(e *entity.Category) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update book set name = ? where name = ?", e.Name, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search books
func (r *CategorySQL) Search(query string) ([]*entity.Category, error) {
	sql_statement := "SELECT * FROM book where title ~* '" + query + "';"
	rows, err := r.db.Query(sql_statement)
	var categories []*entity.Category
	if err != nil {
		panic(err)
	}
	// defer rows.Close()
	var id entity.ID
	var name string
	var updated_at string
	var created_at string
	for rows.Next() {
		switch err := rows.Scan(&id, &name, &created_at, &updated_at); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			categories = append(categories, &entity.Category{ID: id, Name: name})
		default:
			if err != nil {
				return categories, err
			}
		}
	}
	return categories, nil
}

//List books
func (r *CategorySQL) List() ([]*entity.Category, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var books []*entity.Category
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Category
		err = rows.Scan(&b.ID, &b.Name, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

//Delete a book
func (r *CategorySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from category where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (w *CategorySQL) InitialDB() {
	// Drop previous table of same name if one exists.
	_, err := w.db.Exec("DROP TABLE IF EXISTS category;")
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished dropping table (if existed)")
	// Create table.
	_, err = w.db.Exec("CREATE TABLE category (id VARCHAR(50) PRIMARY KEY, name VARCHAR(200), description VARCHAR(1000), updated_at VARCHAR(200), created_at VARCHAR(200));")
	if err != nil {
		panic(err)
	}
	fmt.Print("Finished creating table: Category \n")
}

type ICategoryRepository interface {
	Create(e *entity.Category) (entity.ID, error)
	Get(id entity.ID) (*entity.Category, error)
	Update(e *entity.Category) error
	Search(query string) ([]*entity.Category, error)
	List() ([]*entity.Category, error)
	Delete(id entity.ID) error
	InitialDB()
}
