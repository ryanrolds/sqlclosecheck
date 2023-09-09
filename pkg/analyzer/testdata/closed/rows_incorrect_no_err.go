package closed

import (
	"log"
	"strings"
)

// rowsIncorrectNoErr provides an example of incorrect closing by not calling Err()
func rowsIncorrectNoErr() {
	age := 40
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age) // want "Rows/Stmt/NamedStmt was not closed"
	if err != nil {
		log.Fatal(err)
	}

	names := []string{}

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}

		names = append(names, name)
	}

	// Check for errors from iterating over rows.
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
