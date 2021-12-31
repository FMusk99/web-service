package db

import (
	//"encoding/json"
	//"net/http"
	//"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	//"github.com/rs/cors"
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "123456789"
)

var db *sql.DB // Note the sql package provides the namespace

type User struct {
	ID       int
	Email    string
	Password string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectDb() {
	// Create DB pool
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	fmt.Println(connectionString)
	var err error
	db, err = sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS users;")
	checkError(err)
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE users (id serial PRIMARY KEY, name VARCHAR(50), age INTEGER, language VARCHAR(100), address VARCHAR(200));")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table.
	sql_statement := "INSERT INTO users (name, age, language, address) VALUES ($1, $2, $3, $4);"
	_, err = db.Exec(sql_statement, "jonny", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "jonny1", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "tom", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "jerry", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "z", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	_, err = db.Exec(sql_statement, "", 20, "golang", "34 Hoang Viet, Tan Binh, Ho Chi Minh")
	checkError(err)
	// fmt.Println("Inserted 3 rows of data")
}

func GetUser(keyword string) []UserModel {
	sql_statement := "SELECT * from users where name LIKE '%" + keyword + "%';"
	// fmt.Println(sql_statement)
	rows, err := db.Query(sql_statement)
	// fmt.Println(rows)
	checkError(err)
	defer rows.Close()

	// Read rows from table.
	var id int
	var name string
	var age int
	var language string
	var address string

	// var response []JsonResponse
	var user []UserModel
	for rows.Next() {
		switch err := rows.Scan(&id, &name, &age, &language, &address); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			// fmt.Printf("Data row = (%d, %s, %d)\n", id, name, age)
			user = append(user, UserModel{Id: id, Name: name, Age: age, Language: language, Address: address})
		default:
			checkError(err)
		}
	}
	// fmt.Println(user)
	return user
}

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Language string `json:"language"`
	Address  string `json:"address"`
}
