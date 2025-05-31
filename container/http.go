package container

import (
	"context"
	"github.com/labstack/echo/v4"
	"rbac-project/authenticate"
	"rbac-project/authenticate/file"
	file2 "rbac-project/authorize/file"
	"rbac-project/base"
	"rbac-project/user"
)

func InitServer(ctx context.Context) (*echo.Echo, error) {

	accessManager := file2.NewFileAccessManager()
	authenticator := file.NewFileAuthenticator()

	if err := user.InitController(ctx, accessManager, authenticator); err != nil {
		return nil, err
	}
	authenticate.NewController(authenticator)

	echoObj := echo.New()
	base.InitControllers(echoObj)

	return echoObj, nil
}
