## Set of Extensios for echo framework

[![GoDoc](https://godoc.org/github.com/mhewedy/echox?status.svg)](https://godoc.org/github.com/mhewedy/echox)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhewedy/echox)](https://goreportcard.com/report/github.com/mhewedy/echox)

### `GormAduit` middleware
It associate current user (if exits from JWT token) with current db refernce so when you save an entity, the `CreatedBy`, `UpdatedBy` columns got populated.
```go
// First your entity should embed audited.AuditedModel:
type Product struct {
	gorm.Model
	audited.AuditedModel
	code string
}

// defines the middleware rerferencing the *gorm.DB instance
e.Use(middlewarex.GormAudit(db))

// Then get gorm db reference in the handler method via:
db := middlewarex.GetGormDB(context)    // context is of type echo.Context
// use the db reference to save your entity instances.
```

### `HasRole` middleware
```go
// HasRole middleware, depends on registration of middleware.JWT
e.POST("/create-product", createProduct, middlewarex.HasRole("admin"))

```
