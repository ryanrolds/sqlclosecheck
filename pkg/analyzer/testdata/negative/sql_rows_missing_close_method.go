package negative

import (
	"database/sql"
	"log"
	"strings"
)

func sqlRowsMethodMissingClose() {
	server := BadServer{}
	server.MissingCloseInOtherMethod()
}

type BadServer struct {
	Rows *sql.Rows
}

func (s *BadServer) BadCloseRows() {
	// BAD: s.Rows.Close()
}

func (s *BadServer) MissingCloseInOtherMethod() {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age) // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	s.Rows = rows

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	s.BadCloseRows()

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
