package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	//Component.Setup()
	db, err := sql.Open("sqlite3", "./guestbook.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//db.Exec("create table foo (id integer not null primary key, name text);")

	for {
	}
}
