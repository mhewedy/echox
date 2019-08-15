package xmiddleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type HasRoleConfig struct {
	JWTContextKey string
	RolesClaim    string
}

var DefaultHasRoleConfig = HasRoleConfig{
	JWTContextKey: middleware.DefaultJWTConfig.ContextKey,
	RolesClaim:    "roles",
}

func HasRole(roles ...string) echo.MiddlewareFunc {
	return HasRoleWithConfig(DefaultHasRoleConfig, roles...)
}

func HasRoleWithConfig(config HasRoleConfig, roles ...string) echo.MiddlewareFunc {

	var contains = func(e string, s []interface{}) bool {
		for _, n := range s {
			if e == n {
				return true
			}
		}
		return false
	}
	var httpErr = func(role string) error {
		return &echo.HTTPError{
			Code:    http.StatusForbidden,
			Message: fmt.Sprintf("role(s) [%s] required to access the resource", role)}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get(config.JWTContextKey)
			if user == nil {
				return httpErr(strings.Join(roles, ", "))
			}

			claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
			userRoles, ok := claims[config.RolesClaim].([]interface{})
			if !ok {
				return httpErr(strings.Join(roles, ", "))
			}

			for _, r := range roles {
				if !contains(r, userRoles) {
					return httpErr(r)
				}
			}

			return next(c)
		}
	}
}