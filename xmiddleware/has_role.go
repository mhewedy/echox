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

// HasRole middleware, depends on registration of middleware.JWT
func HasRole(roles ...string) echo.MiddlewareFunc {
	return HasRoleWithConfig(DefaultHasRoleConfig, roles...)
}

// HasRole middleware, depends on registration of middleware.JWT
func HasRoleWithConfig(config HasRoleConfig, roles ...string) echo.MiddlewareFunc {

	var contains = func(e string, s []interface{}) bool {
		for _, n := range s {
			if e == n {
				return true
			}
		}
		return false
	}
	var errForbidden = func(role string) error {
		return &echo.HTTPError{
			Code:    http.StatusForbidden,
			Message: fmt.Sprintf("role(s) [%s] required to access the resource", role)}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get(config.JWTContextKey)
			if user == nil {
				return errForbidden(strings.Join(roles, ", "))
			}

			claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
			userRoles, ok := claims[config.RolesClaim].([]interface{})
			if !ok {
				return errForbidden(strings.Join(roles, ", "))
			}

			for _, r := range roles {
				if !contains(r, userRoles) {
					return errForbidden(r)
				}
			}

			return next(c)
		}
	}
}
