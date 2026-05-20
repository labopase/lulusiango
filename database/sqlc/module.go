package db

import (
	"github.com/halimdotnet/lulusiango/internals/shared/store/pgsql"
	"go.uber.org/fx"
)

// Module is the Fx module that provides the SQLC Querier interface.
var Module = fx.Module("database-sqlc",
	fx.Provide(
		func(c pgsql.Client) Querier {
			return New(c.Pool())
		},
	),
)
