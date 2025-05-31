package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rbac-project/authenticate"
	"rbac-project/authorize"
	"rbac-project/base"
	"rbac-project/middleware"
	"rbac-project/response"
)

const (
	resourceUser string = "user"
	actionGet    string = "get"
	actionPost   string = "post"
)

type controller struct {
	accessManager authorize.AccessManager
	authenticator authenticate.Authenticator
}

func NewController(
	accessManager authorize.AccessManager,
	authenticator authenticate.Authenticator,
) base.Controller {
	return base.AddController(&controller{
		accessManager: accessManager,
		authenticator: authenticator,
	})
}

func (c *controller) RegisterRoutes(router *echo.Echo) {
	v1 := router.Group("/v1")

	users := v1.Group("/users")
	users.Use(middleware.Authenticate(c.authenticator))
	users.Add(http.MethodGet, "/:id", base.EchoHandler(c.getByID), c.authorize(actionGet))
	users.Add(http.MethodPost, "/create", base.EchoHandler(c.create), c.authorize(actionPost))
}

func (c *controller) getByID(*http.Request) (response.BaseResponse, error) {
	return response.BaseResponse{
		Data: "success",
	}, nil
}

func (c *controller) create(r *http.Request) (response.BaseResponse, error) {
	return response.BaseResponse{
		Data: "success create",
	}, nil
}

func (c *controller) authorize(actionType string) echo.MiddlewareFunc {
	return middleware.Authorize(c.accessManager, resourceUser, actionType)
}
