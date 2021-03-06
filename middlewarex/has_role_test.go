package middlewarex

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHasRoleWithTokenHasAdminRole(t *testing.T) {

	jwt := middleware.JWT([]byte("secret"))
	hasRole := jwt(HasRole("admin")(func(context echo.Context) error {
		return nil
	}))
	// token contains roles: [admin]
	context := getContext("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY3NDg3MTksImlkIjoxMDEsInJvbGVzIjpbImFkbWluIiwic3VwZXJ2aXNvciJdfQ.ZxC_XzPvPv6k4BwkvqU7qKoLA7Bz01oi2vzPNWMGba4")
	err := hasRole(context)

	assert.Nil(t, err)
}

func TestHasRoleWithTokenWithoutAdminRole(t *testing.T) {

	jwt := middleware.JWT([]byte("secret"))
	hasRole := jwt(HasRole("admin")(func(context echo.Context) error {
		return nil
	}))
	// token doest NOT contain roles: [admin]
	context := getContext("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY3NjI4MjYsImlkIjoxMDEsInJvbGVzIjpbInVzZXIiXX0.8XcD9EJzh7MuCdfDj7xO3_bI885h13H6ZEmvHeEm1r8")
	err := hasRole(context)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "[admin]")
	assert.Equal(t, err.(*echo.HTTPError).Code, 403)
}

func getContext(token string) echo.Context {
	request, _ := http.NewRequest("", "", strings.NewReader(""))
	request.Header.Set("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()
	context := echo.New().NewContext(request, response)
	return context
}
