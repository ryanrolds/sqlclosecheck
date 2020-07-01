package stmt

import (
	"database/sql"
)

func returnStmt() (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, "SELECT username FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func returnStmtShort() (*sql.Stmt, error) {
	return db.PrepareContext(ctx, "SELECT username FROM users WHERE id = ?")
}
