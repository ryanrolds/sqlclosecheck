package returned

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx context.Context

	db *sql.DB

	pgxTx   pgx.Tx
	pgxConn *pgx.Conn
	pgxPool *pgxpool.Pool
)
