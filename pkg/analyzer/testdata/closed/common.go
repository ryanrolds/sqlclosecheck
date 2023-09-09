package closed

import (
	"context"
	"database/sql"
)

var (
	ctx context.Context
	db  *sql.DB
)
