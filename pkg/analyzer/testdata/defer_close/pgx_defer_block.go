package defer_close

import (
	"log"
)

func pgxDeferBlockTx() {
	rows, err := pgxTx.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}

func pgxDeferBlockConn() {
	rows, err := pgxConn.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}

func pgxDeferBlockPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}
