package sqlx_examples

import (
	"context"
	"github.com/jmoiron/sqlx"
)

var (
	ctx context.Context
	db  *sqlx.DB
)
