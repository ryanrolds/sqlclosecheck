package rows

import (
	"database/sql"
	"log"
)

func returnRows() (*sql.Rows, error) {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	return rows, nil
}

func returnRowsShort() (*sql.Rows, error) {
	age := 27
	return db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
}
