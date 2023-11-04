package negative

import (
	"log"
)

func pgxMissingCloseTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}

func gxMissingCloseConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}

func gxMissingClosePgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users") // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	_ = rows
}
