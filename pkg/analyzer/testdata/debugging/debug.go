package debugging

import (
	"log"
)

func correctDeferBlockPgxPool() {
	rows, err := pgxPool.Query(ctx, "SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		rows.Close()
	}()
}
