package rows

import (
	"log"
	"strings"
)

// nextCorrcet provides an example of corrcet closing of rows without a defer
func nextCorrcet() {
	age := 40
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
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
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
