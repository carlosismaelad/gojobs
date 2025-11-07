package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run The Server
	router.Run(":8080")
}