package pgx_examples

import (
	"log"
)

func correctDeferTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func correctDeferConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
}

func correctDeferPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
}
