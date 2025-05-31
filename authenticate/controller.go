package authenticate

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"rbac-project/base"
	"rbac-project/response"
)

type controller struct {
	authenticator Authenticator
}

func NewController(
	authenticator Authenticator,
) base.Controller {
	return base.AddController(&controller{
		authenticator: authenticator,
	})
}

func (c *controller) RegisterRoutes(router *echo.Echo) {
	v1 := router.Group("/v1")

	v1.Add(http.MethodPost, "/login", base.EchoHandler(c.login))
}

func (c *controller) login(request *http.Request) (response.BaseResponse, error) {
	var req LoginRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return response.BaseResponse{}, err
	}

	token, err := c.authenticator.Login(req)
	if err != nil {
		return response.BaseResponse{}, err
	}

	return response.BaseResponse{
		Data: token,
	}, nil
}
