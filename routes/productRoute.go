package routes

import (
	"sesi-6/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(route *gin.Engine, db *gorm.DB) {
	route.POST("/api/product", func(c *gin.Context) {
		controllers.CreateProduct(c, db)
	})

	route.PUT("/api/product/:id", func(c *gin.Context) {
		controllers.UpdateProduct(c, db)
	})

	route.GET("/api/product/:id", func(c *gin.Context) {
		controllers.GetProductById(c, db)
	})

	route.GET("/api/product/:id/variants", func(c *gin.Context) {
		controllers.GetProductWithVariant(c, db)
	})
}
