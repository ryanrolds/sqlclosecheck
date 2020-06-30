package sqlx_examples

import (
	_ "github.com/go-sql-driver/mysql"
)

func correctDefer() {
	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		return
	}

	defer rows.Close()
}
