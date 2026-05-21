package user

import "github.com/halimdotnet/lulusiango/internals/shared/api/rest"

type Endpoint interface {
	Register(engine rest.Engine)
}
