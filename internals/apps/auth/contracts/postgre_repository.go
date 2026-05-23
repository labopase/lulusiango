package contracts

import (
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/pgsql"
	"go.uber.org/fx"
)

type PostgresqlRepository struct {
	fx.In

	log  logger.Logger
	pool pgsql.Client
	//ProductsGroup   *echo.Group `name:"product-echo-group"`
}
