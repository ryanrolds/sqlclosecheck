package pgx_examples

import (
	"log"

	"github.com/jackc/pgx/v5"
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

func correctDeferBlockWithParameterTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows pgx.Rows) {
		rows.Close()
	}(rows)
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

func correctDeferBlockWithParameterConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows pgx.Rows) {
		rows.Close()
	}(rows)
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

func correctDeferBlockWithParameterPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows pgx.Rows) {
		rows.Close()
	}(rows)
}
