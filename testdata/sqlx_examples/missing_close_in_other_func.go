package sqlx_examples

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func ForgetCloseSqlxStmt(stmt *sqlx.Stmt) {
	if stmt != nil {
		// stmt.Close()
	}
}

func (s Server) DeferForgetCloseInOtherFunc() {
	stmt, err := db.Preparex("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer ForgetCloseSqlxStmt(stmt)

	rows := stmt.QueryRow()
	fmt.Printf("%v", rows)
}
