package closed

import (
	"database/sql"
	"log"
	"strings"
)

func correctMethodClose() {
	server := Server{}
	server.CloseInOtherMethod()
}

type Server struct {
	Rows *sql.Rows
}

func (s *Server) Close() {
	s.Rows.Close()
}

func (s Server) CloseInOtherMethod() {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
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

	s.Close()

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
