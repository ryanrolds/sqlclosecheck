package sqlx_examples

import (
	"github.com/jmoiron/sqlx"
)

func returnRows() (*sqlx.Rows, error) {
	age := 27
	rows, err := db.Queryx("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
