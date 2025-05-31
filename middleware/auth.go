package middleware

import (
	"errors"
	"fmt"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	"net/http"
	"rbac-project/authenticate"
	"rbac-project/authorize"
	"strconv"
	"strings"
)

type userProfile struct {
	ID int64 `json:"id"`
}

const (
	userProfileKey = "UserProfile"
)

func Authorize(accessManager authorize.AccessManager, resource, action string) echo.MiddlewareFunc {
	return casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
		UserGetter: func(c echo.Context) (string, error) {
			uac, ok := c.Get(userProfileKey).(userProfile)
			if !ok {
				fmt.Println("no user profile found")
				return "", errors.New("unauthorized")
			}

			return strconv.FormatInt(uac.ID, 10), nil
		},
		EnforceHandler: func(c echo.Context, user string) (bool, error) {
			allowed, err := accessManager.Check(user, resource, action)
			if err != nil {
				fmt.Println("got err when enforcing: ", err)
				return false, errors.New("server internal error")
			}

			if !allowed {
				return false, errors.New("forbidden")
			}

			return true, nil
		},
	})
}

func Authenticate(authenticator authenticate.Authenticator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationValues := c.Request().Header["Authorization"]
			if len(authorizationValues) == 0 {
				fmt.Println("no header found")
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			authToken := strings.Split(authorizationValues[0], " ")[1]

			userID, err := authenticator.Authenticate(authToken)
			if err != nil {
				fmt.Println(err)
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			c.Set(userProfileKey, userProfile{
				ID: userID,
			})

			return next(c)
		}
	}
}
