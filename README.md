
[![GoDoc](https://godoc.org/github.com/mhewedy/echox?status.svg)](https://godoc.org/github.com/mhewedy/echox)

`GormAduit` middleware
```go
// GormAudit middleware, depends on registration of middleware.JWT
// Allow injecting CreatedBy and UpdatedBy into the db fields
e.Use(middlewarex.GormAudit(db))

// Then get gorm db reference in the handler method via:
db := middlewarex.GetGormDB(context)    // context is of type echo.Context
```

`HasRole` middleware
```go
// HasRole middleware, depends on registration of middleware.JWT
e.POST("/create-product", createProduct, middlewarex.HasRole("admin"))

```
