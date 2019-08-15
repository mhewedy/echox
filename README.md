`GormAduit` middleware
```go
// GormAudit middleware, depends on registration of middleware.JWT
// Allow injecting CreatedBy and UpdatedBy into the db fields
e.Use(xmiddleware.GormAudit(db))

// Then get gorm db reference in the handler method via:
// context is of type echo.Context
db := xmiddleware.GetGormDB(context)
```

`HasRole` middleware
```go
// HasRole middleware, depends on registration of middleware.JWT
e.POST("/create-product", createProduct, xmiddleware.HasRole("admin"))

```
