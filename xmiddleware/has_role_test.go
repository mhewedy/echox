package xmiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	context := getContext("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY3NDg3MTksImlkIjoxMDEsInJvbGVzIjpbImFkbWluIiwic3VwZXJ2aXNvciJdfQ.ZxC_XzPvPv6k4BwkvqU7qKoLA7Bz01oi2vzPNWMGba4")
	err := hasRole(context)

	if err != nil {
		t.Errorf("HasRole() should not return error")
	}
}

func TestHasRoleWithTokenWithoutAdminRole(t *testing.T) {

	jwt := middleware.JWT([]byte("secret"))
	hasRole := jwt(HasRole("admin")(func(context echo.Context) error {
		return nil
	}))

	context := getContext("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY3NjI4MjYsImlkIjoxMDEsInJvbGVzIjpbInVzZXIiXX0.8XcD9EJzh7MuCdfDj7xO3_bI885h13H6ZEmvHeEm1r8")
	err := hasRole(context)

	if err == nil {
		t.Errorf("HasRole() should return error")
	}
	strings.Contains(err.Error(), "[admin]")
}

func getContext(token string) echo.Context {
	request, _ := http.NewRequest("", "", strings.NewReader(""))
	request.Header.Set("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()
	context := echo.New().NewContext(request, response)
	return context
}
