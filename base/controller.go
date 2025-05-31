package base

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rbac-project/response"
)

var controllers []Controller

type Controller interface {
	RegisterRoutes(e *echo.Echo)
}

func InitControllers(e *echo.Echo) {
	for _, controller := range controllers {
		controller.RegisterRoutes(e)
	}
}

func AddController(controller Controller) Controller {
	controllers = append(controllers, controller)
	return controller
}

func EchoHandler(fn func(r *http.Request) (response.BaseResponse, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := fn(c.Request())
		if err != nil {
			return err
		}

		resp.Status = http.StatusOK
		resp.Message = http.StatusText(resp.Status)

		return c.JSON(http.StatusOK, resp)
	}
}
