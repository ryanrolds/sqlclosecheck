package sqlx_examples

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	Stmt *sqlx.Stmt
}

func (s *Server) Close() {
	s.Stmt.Close()
}

func (s Server) CloseInOtherFunc() {
	stmt, err := db.Preparex("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	s.Stmt = stmt

	rows := stmt.QueryRow()
	fmt.Printf("%v", rows)

	s.Close()
}
