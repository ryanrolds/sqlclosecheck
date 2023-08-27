package pgx_examples

import (
	"github.com/jackc/pgx/v5"
	"log"
)

func returnStmtTx() (pgx.Rows, error) {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

func returnStmtShortTx() (pgx.Rows, error) {
	return pgxTx.Query(ctx, "SELECT username FROM users")
}

func returnStmtConn() (pgx.Rows, error) {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

func returnStmtShortConn() (pgx.Rows, error) {
	return pgxConn.Query(ctx, "SELECT username FROM users")
}

func returnStmtPgxPool() (pgx.Rows, error) {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return rows, nil
}

func returnStmtShortPgxPool() (pgx.Rows, error) {
	return pgxPool.Query(ctx, "SELECT username FROM users")
}
