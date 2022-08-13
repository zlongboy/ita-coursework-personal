package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() {
	dbUser := os.Getenv("BOOKS_DB_USER")
	dbPass := os.Getenv("BOOKS_DB_SECRET")
	hostPort := "127.0.0.1:3306"
	dbName := "books_v2"

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", dbUser, dbPass, hostPort, dbName))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// TODO:
	// Write update books query
	// Write update publishers query
	// Write update authors query
	// Write fetchAll (for authorId) query

	// *** TESTING DB CONNECTION - TO DELETE**
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
