package middlewarex

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/qor/audited"
)

const GormDBKey = "Gorm_DB"

type GormAuditedConfig struct {
	JWTContextKey string
	IDClaim       string
}

var DefaultGormAuditedConfig = GormAuditedConfig{
	JWTContextKey: middleware.DefaultJWTConfig.ContextKey,
	IDClaim:       "id",
}

// GormAudit middleware, depends on registration of middleware.JWT
// Allow injecting CreatedBy and UpdatedBy into the db fields
func GormAudit(db *gorm.DB) echo.MiddlewareFunc {
	return GormAuditWithConfig(db, DefaultGormAuditedConfig)
}

// GormAudit middleware, depends on registration of middleware.JWT
// Allow injecting CreatedBy and UpdatedBy into the db fields
func GormAuditWithConfig(db *gorm.DB, config GormAuditedConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if user := c.Get(config.JWTContextKey); user != nil {

				claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
				currentUser := struct{ ID uint }{
					ID: uint(claims[config.IDClaim].(float64)),
				}

				db = db.Set("audited:current_user", currentUser)
				c.Set(GormDBKey, db)
			}
			return next(c)
		}
	}
}

func GetGormDB(c echo.Context) *gorm.DB {
	return c.Get(GormDBKey).(*gorm.DB)
}
