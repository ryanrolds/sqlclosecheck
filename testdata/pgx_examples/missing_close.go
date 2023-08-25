package pgx_examples

import (
	"log"
)

func missingCloseTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}

func missingCloseConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}

func missingClosePgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}
