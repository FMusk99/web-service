package config

const (
	DB_USER                = "postgres"
	DB_PASSWORD            = "123456789"
	DB_DATABASE            = "postgres"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)

// const (
// 	// Initialize connection constants.
// 	HOST     = "localhost"
// 	DATABASE = "postgres"
// 	USER     = "postgres"
// 	PASSWORD = "123456789"
// )
// var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

// var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
// fmt.Println(connectionString)
// var err error
// db, err = sql.Open("postgres", connectionString)
