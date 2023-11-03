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

func (s Server) CloseInOtherMethod() {
	stmt, err := db.Preparex("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	s.Stmt = stmt

	rows := stmt.QueryRow()
	fmt.Printf("%v", rows)

	s.Close()
}

func CloseSqlxStmt(stmt *sqlx.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

func (s Server) DeferCloseInOtherFunc() {
	stmt, err := db.Preparex("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer CloseSqlxStmt(stmt)

	rows := stmt.QueryRow()
	fmt.Printf("%v", rows)
}
