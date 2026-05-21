package users

import (
	createuser "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/features/create_user/v1"
	deleteuser "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/features/delete_user/v1"
	getuserbyemail "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/features/get_user_by_email/v1"
	getuserbyid "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/features/get_user_by_id/v1"
	updateuser "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/features/update_user/v1"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"go.uber.org/fx"
)

var Module = fx.Module("user_fx",
	fx.Provide(
		repository.NewPostgreRepository,
	),
	fx.Provide(
		getuserbyid.NewGetUserByID,
		getuserbyemail.NewGetUserByEmail,
	),
	fx.Provide(
		createuser.NewCreateUser,
		updateuser.NewUpdateUser,
		deleteuser.NewDeleteUser,
	),
	fx.Provide(
		fx.Annotate(
			getuserbyid.NewGetUserByIDEndpoints,
			fx.As(new(user.Endpoint)),
			fx.ResultTags(`group:"user_endpoints"`),
		),
		fx.Annotate(
			getuserbyemail.NewGetUserByEmailEndpoints,
			fx.As(new(user.Endpoint)),
			fx.ResultTags(`group:"user_endpoints"`),
		),
		fx.Annotate(
			createuser.NewCreateUserEndpoints,
			fx.As(new(user.Endpoint)),
			fx.ResultTags(`group:"user_endpoints"`),
		),
		fx.Annotate(
			updateuser.NewUpdateUserEndpoints,
			fx.As(new(user.Endpoint)),
			fx.ResultTags(`group:"user_endpoints"`),
		),
		fx.Annotate(
			deleteuser.NewDeleteUserEndpoints,
			fx.As(new(user.Endpoint)),
			fx.ResultTags(`group:"user_endpoints"`),
		),
	),
	fx.Invoke(
		fx.Annotate(
			registerEndpoints,
			fx.ParamTags("", `group:"user_endpoints"`),
		),
	),
)

func registerEndpoints(engine rest.Engine, endpoints []user.Endpoint) {
	for _, ep := range endpoints {
		ep.Register(engine)
	}
}
