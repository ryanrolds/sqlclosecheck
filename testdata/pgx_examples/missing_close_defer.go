package pgx_examples

import (
	"log"

	"github.com/jackc/pgx/v5"
)

func missingCloseDeferBlock() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	_ = rows

	defer func() {

	}()
}

func missingCloseDeferBlockWithParameter() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	_ = rows

	defer func(rows pgx.Rows) {

	}(rows)
}
