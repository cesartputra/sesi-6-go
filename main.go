package main

import (
	"sesi-6/database"
	"sesi-6/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	r := gin.Default()

	routes.ProductRoutes(r, db)
	routes.VariantRoutes(r, db)

	r.Run(":8080")
}
