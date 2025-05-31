package user

import (
	"context"
	"rbac-project/authenticate"
	"rbac-project/authorize"
)

func InitController(
	_ context.Context,
	am authorize.AccessManager,
	authenticator authenticate.Authenticator,
) error {
	NewController(am, authenticator)
	return nil
}
