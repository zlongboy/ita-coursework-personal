package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(bv []BookValues, autIDs, pubIDs map[string]string) {
	dbUser := os.Getenv("BOOKS_DB_USER")
	dbPass := os.Getenv("BOOKS_DB_SECRET")
	hostPort := "127.0.0.1:3306"
	dbName := "books_v2"

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", dbUser, dbPass, hostPort, dbName))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// BOOKS QUERY
	booksResult, err := db.Exec(booksQuery(bv))
	if err != nil {
		fmt.Println(err)
	}

	bra, err := booksResult.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Books added: %v \n", bra)

	// AUTHORS QUERY
	authResult, err := db.Exec(authorPublisherQuery(autIDs, "authors"))
	if err != nil {
		fmt.Println(err)
	}

	ara, err := authResult.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Authors added: %v \n", ara)

	// PUBLISHERS QUERY
	pubResult, err := db.Exec(authorPublisherQuery(pubIDs, "publishers"))
	if err != nil {
		fmt.Println(err)
	}

	pra, err := pubResult.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Publishers added: %v \n", pra)

	// BOOKSAUTHORS QUERY
	_, err := db.Exec(indexTableQuery(bv, autIDs, "booksauthors")) // TODO: Don't need to the result from db.Exec
	if err != nil {
		fmt.Println(err)
	}

	// BOOKSPUBLISHER QUERY
	_, err := db.Exec(indexTableQuery(bv, pubIDs, "bookspublishers")) // TODO: Don't need to the result from db.Exec
	if err != nil {
		fmt.Println(err)
	}
}

// QUERY BUILDERS
func booksQuery(bv []BookValues) string {
	// var valueStrings, valueArgs []string
	var valueArgs []string
	for _, b := range bv {
		// valueStrings = append(valueStrings, "(?, ?, ?, ?)") // TODO: Use placeholders to prevent SQL injection
		arg := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", b.ID, escapeSQL(b.Title), escapeSQL(b.Subtitle), escapeSQL(b.Desc), b.PublishDate, b.Country, b.ImageURL, b.PurchaseURL)
		valueArgs = append(valueArgs, arg)
		// TODO: Handle inserting different types for row.Price (float), row.PDF (bool)
	}
	bq := fmt.Sprintf("INSERT INTO books (id, title, subtitle, blurb, publish_date, country, image_url, purchase_url) VALUES %s", strings.Join(valueArgs, ","))
	// fa := strings.Join(valueArgs, ",")
	return bq
}

func indexTableQuery(bv []BookValues, m map[string]string, table string) string { // Constructs booksauthors, bookspublishers query string

	var valueArgs []string
	for _, b := range bv {
		var tname string
		if table == "booksauthors" {
			tname = b.Author
		}
		if table == "bookspublishers" {
			tname = b.Publisher
		}

		arg := fmt.Sprintf("('%s','%s')", b.ID, m[tname])
		valueArgs = append(valueArgs, arg)
	}
	iq := fmt.Sprintf("INSERT INTO %s VALUES %s", table, strings.Join(valueArgs, ","))
	fmt.Println(iq)
	return iq
}

func authorPublisherQuery(m map[string]string, table string) string {

	var valueArgs []string
	for k, id := range m {
		arg := fmt.Sprintf("('%s','%s')", id, k)
		valueArgs = append(valueArgs, arg)
	}
	apq := fmt.Sprintf("INSERT INTO %s VALUES %s", table, strings.Join(valueArgs, ","))
	return apq
}

// UTILS
func escapeSQL(t string) string {
	return strings.Replace(t, "'", "''", -1)
}
