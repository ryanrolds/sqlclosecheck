package sqlx_examples

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

func missingCloseInBlock() {
	age := 27
	rows, err := db.Queryx("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// rows.Close()
	}()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}

func missingCloseInBlockWithParameter() {
	age := 27
	rows, err := db.Queryx("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sqlx.Rows) {
		// rows.Close()
	}(rows)

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
