package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Practice")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3000)/dbname") // user:password

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM authors")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var author Author

		err = results.Scan(author.Name) // &author.Name
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(author.Name)
	}
}
