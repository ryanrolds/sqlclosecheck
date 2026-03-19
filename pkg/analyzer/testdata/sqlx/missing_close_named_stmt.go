package sqlx

import (
	"log"
)

func missingCloseNamedStmt() {
	stmt, err := db.PrepareNamed("SELECT * FROM users WHERE id = :id") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	// defer stmt.Close()

	_ = stmt // No need to use stmt
}
