package defer_close

import (
	"log"
)

func pgxDeferTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func pgxDeferConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
}

func pgxDeferPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
}
