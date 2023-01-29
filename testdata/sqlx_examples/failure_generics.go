package sqlx_examples

import "database/sql"

func SqlCloseCheck(db *sql.DB, a int) {
	rows, _ := db.Query("select id from tb")
	for rows.Next() {

	}
}

func SqlCloseCheckG[T ~int64](db *sql.DB, a T) {
	rows, _ := db.Query("select id from tb")
	for rows.Next() {

	}
}
