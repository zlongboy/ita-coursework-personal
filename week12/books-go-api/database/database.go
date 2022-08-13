package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	// "github.com/zlongboy/ita-coursework-personal/week12/books-api/util"
)

// TODO: Split out to util package (repeated code in client package)
func getEnvMap() map[string]string {

	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error reading .env to map")
	}
	return envMap
}

func main() {
	dbName := "books_v2"
	dbUser := getEnvMap()["BOOKS_DB_USER"]
	dbPass := getEnvMap()["BOOKS_DB_SECRET"]
	hostPort := "127.0.0.1:3306"

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", dbUser, dbPass, hostPort, dbName))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM publishers")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var publisher Publisher

		err = results.Scan(&publisher.PublisherName, &publisher.ID) // Tells scan method where to store the data
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(publisher.PublisherName, publisher.ID)
	}
}
