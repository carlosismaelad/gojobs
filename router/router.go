package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "API ok!",
    })
  })
	router.Run(":8080")
}