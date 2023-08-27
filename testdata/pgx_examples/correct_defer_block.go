package pgx_examples

import (
	"log"
)

func correctDeferBlockTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}

func correctDeferBlockConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}

func correctDeferBlockPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}
