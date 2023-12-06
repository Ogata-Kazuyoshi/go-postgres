package main

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=user dbname=go_template sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.Select("*").From("persons").Where(sq.Eq{"age":32}).ToSql()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int
		err = rows.Scan(&name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Name:", name, "Age:", age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}