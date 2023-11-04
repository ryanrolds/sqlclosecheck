package passed

import (
	"database/sql"
	"log"
)

func sqlRowsPassedClosed() {
	rows, err := db.QueryContext(ctx, "SELECT name FROM users")
	if err != nil {
		log.Fatal(err)
	}

	closedPassed(rows)
}

func closedPassed(rows *sql.Rows) {
	rows.Close()
}

func sqlRowsPassedMissingClosed(rows *sql.Rows) {
	rows, err := db.QueryContext(ctx, "SELECT name FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	dontClosedPassed(rows)
}

func dontClosedPassed(*sql.Rows) {

}
