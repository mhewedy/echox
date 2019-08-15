```go

// GormAudit middleware, depends on registration of middleware.JWT
e.Use(xmiddleware.GormAudit(db))

// HasRole middleware, depends on registration of middleware.JWT
e.POST("/create-product", createProduct, xmiddleware.HasRole("admin"))

```
