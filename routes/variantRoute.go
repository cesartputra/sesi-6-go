package routes

import (
	"sesi-6/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VariantRoutes(route *gin.Engine, db *gorm.DB) {
	route.POST("/api/variant", func(c *gin.Context) {
		controllers.CreateVariant(c, db)
	})

	route.PUT("/api/variant/:id", func(c *gin.Context) {
		controllers.UpdateVariantById(c, db)
	})

	route.DELETE("/api/variant/:id", func(c *gin.Context) {
		controllers.DeleteVariantById(c, db)
	})
}
