package infrastructure

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type (
	queryRunner interface {
		sqlx.ExtContext
		sqlx.Ext
		GetContext(
			ctx context.Context,
			dest interface{},
			query string,
			args ...interface{},
		) error
		SelectContext(
			ctx context.Context,
			dest interface{},
			query string,
			args ...interface{},
		) error
	}
)
