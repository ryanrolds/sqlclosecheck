package sqlx_examples

import (
	_ "github.com/go-sql-driver/mysql"
)

func nonDeferClose() {
	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		return
	}

	rows.Close()
}
